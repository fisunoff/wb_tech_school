package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

const DeadOrdersQueue = "dead-orders"

type dlqProducer struct {
	writer kafka.Writer
}

func newDlqProducer(brokers []string) *dlqProducer {
	return &dlqProducer{
		writer: kafka.Writer{
			Addr:                   kafka.TCP(brokers...),
			Topic:                  DeadOrdersQueue,
			Balancer:               &kafka.Hash{},
			AllowAutoTopicCreation: true,
			Async:                  true,
			ErrorLogger:            kafka.LoggerFunc(log.Printf),
		},
	}
}

// HandleMessage отправляет сообщение в DLQ
func (p *dlqProducer) HandleMessage(ctx context.Context, msg []byte) error {
	kafkaMsg := kafka.Message{
		Key:   msg, // не факт, что можно что-то достать, поэтому в ключ кладем сообщение целиком
		Value: msg,
	}
	if err := p.writer.WriteMessages(ctx, kafkaMsg); err != nil {
		log.Printf("[producer] Ошибка записи в DLQ в Kafka: %v", err)
		return err
	}
	return nil
}
