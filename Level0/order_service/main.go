package main

import (
	"fmt"
	"log"
	"order_service/storage"
	"os"
	"path/filepath"
)

func readData(filename string) []byte {
	path := filepath.Join("model/testdata", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}

func main() {
	log.Println("Запускаем сервис...")

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

	//correctData := readData("valid_order.json")
	//data, err := model.SerializeOrder([]byte(correctData))
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//err = db.SaveOrder(&data)
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}

	data, err := db.GetOrderByUID(db.GetDb(), "b563feb7b2b84b6test")
	fmt.Println(data)

	log.Println("Успешно подключились к базе данных!")
}
