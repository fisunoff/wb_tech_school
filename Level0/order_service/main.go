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

	httpSwagger "github.com/swaggo/http-swagger"

	_ "order_service/docs"
	"order_service/kafka"
	"order_service/model"
)

const PORT = ":8080"

// @title           Order Service API
// @version         1.0
// @description     Это API для сервиса заказов.

// @host      localhost:8080
// @BasePath  /
func main() {
	log.Println("Запускаем сервис...")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	dbURL := utils.Env("DB_URL", "")

	maxCacheSizeString := utils.Env("MAX_CACHE_SIZE", "5000")
	maxCacheSize := utils.MustAtoi(maxCacheSizeString)

	startCacheSizeString := utils.Env("START_CACHE_SIZE", "1000")
	startCacheSize := utils.MustAtoi(startCacheSizeString)

	db, err := storage.New(dbURL, maxCacheSize, startCacheSize)
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

	// docs
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	// frontend
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/*", http.StripPrefix("/", fs))

	err := http.ListenAndServe(PORT, router)
	if err != nil {
		log.Fatalf("не удалось запустить сервер: %v", err)
	}
}

func flyProducer(ctx context.Context, brokers []string, topic string, ratePerSec int) {
	generator := func() (key []byte, value []byte, ts time.Time, err error) {
		order, err := model.NewFakeOrder()
		if err != nil {
			return nil, nil, time.Now(), err
		}
		jsonOrder, _ := json.Marshal(order)
		if err != nil {
			return nil, nil, time.Now(), err
		}
		return []byte(order.OrderUID), jsonOrder, order.DateCreated, nil
	}

	go kafka.StartProducing(ctx, brokers, topic, ratePerSec, generator)
}

func flyConsumer(ctx context.Context, brokers []string, topic string, db *storage.Storage) {
	go kafka.StartConsuming(
		ctx,
		kafka.NewReaderConfig(
			brokers,
			topic,
		),
		db,
		brokers,
	)
}
