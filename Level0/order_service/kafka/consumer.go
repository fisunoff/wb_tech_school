package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const COMMIT_ON_ERROR = false // коммитить не надо, если ошибка
const ORDER_CONSUMER = "order-service-consumer"

// Handler — функция обработки сообщения.
type Handler func(ctx context.Context, key, value []byte, ts time.Time) error

func NewReaderConfig(Brokers []string, Topic string) *kafka.ReaderConfig {
	return &kafka.ReaderConfig{
		Brokers:  Brokers,
		Topic:    Topic,
		GroupID:  ORDER_CONSUMER,
		MinBytes: 1 << 10,  // 1KB
		MaxBytes: 10 << 20, // 10MB
		MaxWait:  250 * time.Millisecond,
	}
}

// StartConsuming - получать сообщения из топика и обрабатывать их через handler
func StartConsuming(ctx context.Context, cfg *kafka.ReaderConfig, handler Handler) {
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
			log.Printf("[consumer] Ошибка при получении сообщения: %v", err.Error())
			return
		}
		if err := handler(ctx, msg.Key, msg.Value, msg.Time); err != nil {
			log.Printf("[consumer] Ошибка обработчика: %v", err)
			if COMMIT_ON_ERROR {
				if cerr := reader.CommitMessages(ctx, msg); cerr != nil && ctx.Err() == nil {
					log.Printf("[consumer] Ошибка после коммита: %v", cerr)
				}
			}
			continue
		}

		if err := reader.CommitMessages(ctx, msg); err != nil && ctx.Err() == nil {
			log.Printf("[consumer] Ошибка при коммите: %v", err)
		}
	}

}
