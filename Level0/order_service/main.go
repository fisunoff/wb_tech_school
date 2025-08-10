package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"order_service/utils"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"order_service/handlers"
	"order_service/storage"

	"order_service/model"

	"order_service/kafka"
)

const PORT = ":8080"

func main() {
	log.Println("Запускаем сервис...")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	dbURL := "postgres://myuser:mypassword@db:5432/order_service_db?sslmode=disable"

	db, err := storage.New(dbURL)
	if err != nil {
		log.Fatalf("Не удалось инициализировать БД: %v", err)
	}

	defer func(db *storage.Storage) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Не удалось закрыть подключение к БД: %v", err)
		}
	}(db)

	go flyHttpServer(db)

	enableProducer := utils.Env("PRODUCER_ENABLED", "false") == "true"
	brokers := utils.SplitAndTrim(utils.Env("KAFKA_BROKERS", "kafka:9092"))
	topic := utils.Env("KAFKA_TOPIC", "orders")
	if enableProducer {
		ratePerSec := utils.MustAtoi(utils.Env("PRODUCER_RATE_PER_SECOND", "20"))
		flyProducer(ctx, brokers, topic, ratePerSec)
		log.Printf("Kafka producer: brokers=%v topic=%s rate=%d msg/s", brokers, topic, ratePerSec)
	} else {
		log.Print("Kafka producer выключен")
	}

	enableConsumer := utils.Env("CONSUMER_ENABLED", "false") == "true"
	if enableConsumer {
		flyConsumer(ctx, brokers, topic, db)
		log.Printf("Kafka consumer: brokers=%v topic=%s", brokers, topic)
	} else {
		log.Print("Kafka consumer выключен")
	}

	<-ctx.Done()
}

func flyHttpServer(db *storage.Storage) {
	orderHandler := &handlers.OrderHandler{
		DB: db,
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/order/{order_uid}", orderHandler.GetByUID)

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/*", http.StripPrefix("/", fs))

	err := http.ListenAndServe(PORT, router)
	if err != nil {
		log.Fatalf("не удалось запустить сервер: %v", err)
	}
}

func flyProducer(ctx context.Context, brokers []string, topic string, ratePerSec int) {
	generator := func() (key []byte, value []byte, ts time.Time) {
		order, _ := model.NewFakeOrder()
		jsonOrder, _ := json.Marshal(order)
		return []byte(order.OrderUID), jsonOrder, order.DateCreated
	}

	go kafka.StartProducing(ctx, brokers, topic, ratePerSec, generator)
}

func flyConsumer(ctx context.Context, brokers []string, topic string, db *storage.Storage) {
	generator := func(ctx context.Context, key, value []byte, ts time.Time, passOnDecodeError bool) error {
		order, err := model.SerializeOrder(value)
		if err != nil {
			if passOnDecodeError { // не получилось - ну и не надо
				log.Println("Пропустили некорректную запись из Kafka")
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

	go kafka.StartConsuming(
		ctx,
		kafka.NewReaderConfig(
			brokers,
			topic,
		),
		generator,
	)
}
