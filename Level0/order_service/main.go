package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"order_service/handlers"
	"order_service/storage"
)

const PORT = ":8080"

func main() {
	log.Println("Запускаем сервис...")

	dbURL := "postgres://myuser:mypassword@localhost:54320/order_service_db?sslmode=disable"

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

	err = http.ListenAndServe(PORT, r)
	if err != nil {
		log.Fatalf("не удалось запустить сервер: %v", err)
	}

	log.Printf("Сервер запущен на порту %s", PORT)

}
