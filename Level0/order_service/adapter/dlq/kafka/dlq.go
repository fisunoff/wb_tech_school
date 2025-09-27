package kafka_dlq

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"order_service/internal/repository"
)

const DeadOrdersQueue = "dead-orders"

type KafkaDLQ struct {
	writer *kafka.Writer
}

func NewKafkaDLQ(brokers []string) repository.DeadLetterQueue {
	return &KafkaDLQ{
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(brokers...),
			Topic:                  DeadOrdersQueue,
			Balancer:               &kafka.Hash{},
			AllowAutoTopicCreation: true,
			Async:                  true,
			ErrorLogger:            kafka.LoggerFunc(log.Printf),
		},
	}
}

func (p *KafkaDLQ) Send(ctx context.Context, msg []byte) error {
	kafkaMsg := kafka.Message{
		Key:   msg,
		Value: msg,
	}
	if err := p.writer.WriteMessages(ctx, kafkaMsg); err != nil {
		log.Printf("[dlq] Ошибка записи в DLQ: %v", err)
		return err
	}
	log.Printf("[dlq] Запись отправлена в DLQ")
	return nil
}

func (p *KafkaDLQ) Close() error {
	return p.writer.Close()
}
