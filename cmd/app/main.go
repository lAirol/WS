package main

import (
	"WS/internal/db"
	"WS/internal/route"
	"fmt"
	"net/http"
)

func main() {
	// Инициализация базы данных
	db.InitDB()
	// Маршрутизация
	route.InitRoute()

	fmt.Println("Сервер запущен на http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
