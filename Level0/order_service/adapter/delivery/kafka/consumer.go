package kafka

import (
	"context"
	"fmt"
	"log"
	"order_service/internal/repository"
	"order_service/internal/usecase"
	"time"

	"github.com/segmentio/kafka-go"
)

const CommitOnError = false // коммитить не надо, если ошибка
const OrderConsumer = "order-service-consumer"

// Handler — функция обработки сообщения.
// passOnDecodeError - пропускать при ошибке преобразования в структуру (не валить ошибку)
type Handler func(ctx context.Context, key, value []byte, ts time.Time, passOnDecodeError bool) error

func NewReaderConfig(Brokers []string, Topic string) *kafka.ReaderConfig {
	return &kafka.ReaderConfig{
		Brokers:  Brokers,
		Topic:    Topic,
		GroupID:  OrderConsumer,
		MinBytes: 1 << 10,  // 1KB
		MaxBytes: 10 << 20, // 10MB
		MaxWait:  250 * time.Millisecond,
	}
}

// StartConsuming - получать сообщения из топика и обрабатывать их
func StartConsuming(
	ctx context.Context,
	cfg *kafka.ReaderConfig,
	useCase *usecase.OrderCreateUseCase,
	dlqHandler repository.DeadLetterQueue,
) {
	reader := kafka.NewReader(*cfg)
	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("[consumer] Ошибка при закрытии: %v", err)
		}
	}()

	log.Printf("[consumer]: brokers=%v topic=%s group=%s", cfg.Brokers, cfg.Topic, cfg.GroupID)

	for {
		msg, err := reader.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			log.Printf("[consumer] Ошибка при получении сообщения: %v", err)
			return
		}

		err = handleMessage(ctx, msg.Key, msg.Value, msg.Time, true, useCase, dlqHandler)
		if err != nil {
			log.Printf("[consumer] Ошибка обработки сообщения: %v", err)
			if !CommitOnError {
				continue
			}
		}

		if err = reader.CommitMessages(ctx, msg); err != nil && ctx.Err() == nil {
			log.Printf("[consumer] Ошибка при коммите: %v", err)
		}
	}
}

func handleMessage(
	ctx context.Context,
	key, value []byte,
	ts time.Time,
	passOnDecodeError bool,
	useCase *usecase.OrderCreateUseCase,
	dlqHandler repository.DeadLetterQueue,
) error {
	err := useCase.CreateOrderFromRaw(ctx, value)
	if err != nil {
		if passOnDecodeError { // не получилось - перемещаем в DLQ
			if dlqErr := dlqHandler.Send(ctx, value); dlqErr != nil {
				return fmt.Errorf("не удалось отправить в DLQ: %w", dlqErr)
			}
			return nil
		}
		return err
	}

	log.Printf("[consumer] заказ сохранён, key=%s, ts=%s", string(key), ts)
	return nil
}
