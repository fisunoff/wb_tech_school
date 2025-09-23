package kafka

import (
	"context"
	"log"
	"order_service/interfaces"
	"order_service/model"
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

// StartConsuming - получать сообщения из топика и обрабатывать их через handler
func StartConsuming(
	ctx context.Context,
	cfg *kafka.ReaderConfig,
	db interfaces.DatabaseInterface,
	dlqBrokers []string,
) {
	reader := kafka.NewReader(*cfg)
	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("[consumer] Ошибка при закрытии: %v", err)
		}
	}()

	log.Printf("[consumer]: brokers=%v topic=%s group=%s", cfg.Brokers, cfg.Topic, cfg.GroupID)

	dlqHandler := newDlqProducer(dlqBrokers)

	for {
		msg, err := reader.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			log.Printf("[consumer] Ошибка при получении сообщения: %v", err.Error())
			return
		}
		if err = baseHandler(ctx, msg.Key, msg.Value, msg.Time, true, db, dlqHandler); err != nil {
			log.Printf("[consumer] Ошибка обработчика: %v", err)
			if CommitOnError {
				if cerr := reader.CommitMessages(ctx, msg); cerr != nil && ctx.Err() == nil {
					log.Printf("[consumer] Ошибка после коммита: %v", cerr)
				}
			}
			continue
		}

		if err = reader.CommitMessages(ctx, msg); err != nil && ctx.Err() == nil {
			log.Printf("[consumer] Ошибка при коммите: %v", err)
		}
	}

}

func baseHandler(
	ctx context.Context,
	key, value []byte,
	ts time.Time,
	passOnDecodeError bool,
	db interfaces.DatabaseInterface,
	dlqHandler *dlqProducer,
) error {
	order, err := model.SerializeOrder(value)
	if err != nil {
		if passOnDecodeError { // не получилось - перемещаем в DLQ
			err = dlqHandler.HandleMessage(ctx, value)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	err = db.SaveOrder(&order)
	if err != nil {
		return err
	}
	log.Printf("[consumer] получили order uid=%s key=%s ts=%s", order.OrderUID, string(key), ts)
	return nil
}
