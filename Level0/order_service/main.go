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

	gen "order_service/model"

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

	orderHandler := &handlers.OrderHandler{
		DB: db,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/order/{order_uid}", orderHandler.GetByUID)

	enableProducer := utils.Env("PRODUCER_ENABLED", "false") == "true"
	if enableProducer {
		brokers := utils.SplitAndTrim(utils.Env("KAFKA_BROKERS", "kafka:9092"))
		topic := utils.Env("KAFKA_TOPIC", "orders")
		ratePerSec := utils.MustAtoi(utils.Env("PRODUCER_RATE_PER_SECOND", "20"))

		generator := func() (key []byte, value []byte, ts time.Time) {
			order, _ := gen.NewFakeOrder()
			jsonOrder, _ := json.Marshal(order)
			return []byte(order.OrderUID), jsonOrder, order.DateCreated
		}

		go kafka.StartProducing(ctx, brokers, topic, ratePerSec, generator)
		log.Printf("Kafka producer: brokers=%v topic=%s rate=%d msg/s", brokers, topic, ratePerSec)
	} else {
		log.Print("Kafka producer выключен")
	}

	err = http.ListenAndServe(PORT, r)
	if err != nil {
		log.Fatalf("не удалось запустить сервер: %v", err)
	}

	log.Printf("Сервер запущен на порту %s", PORT)

}
