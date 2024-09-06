package main

import (
	"WS/system/db"
	"WS/system/handlers"
	"fmt"
	"net/http"
)

func main() {
	// Инициализация базы данных
	db.InitDB()
	// Маршрутизация
	handlers.InitRoute()

	fmt.Println("Сервер запущен на http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
