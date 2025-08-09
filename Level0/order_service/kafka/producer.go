package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// Generator должен вернуть ключ (для партиционирования), полезную нагрузку и timestamp сообщения.
type Generator func() (key []byte, value []byte, ts time.Time)

// StartProducing запускает цикл отправки сообщений с частотой rate сообщений в секунду.
func StartProducing(ctx context.Context, brokers []string, topic string, rate int, gen Generator) {
	interval := time.Second / time.Duration(rate)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	writer := &kafka.Writer{
		Addr:                   kafka.TCP(brokers...),
		Topic:                  topic,
		Balancer:               &kafka.Hash{},
		AllowAutoTopicCreation: true,
	}
	defer func() {
		if err := writer.Close(); err != nil {
			log.Printf("[producer] Ошибка при закрытии: %v", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			key, val, ts := gen()
			msg := kafka.Message{
				Key:   key,
				Value: val,
				Time:  ts,
			}
			if err := writer.WriteMessages(ctx, msg); err != nil {
				log.Printf("[producer] Ошибка записи в Kafka: %v", err)
			}
		}
	}
}
