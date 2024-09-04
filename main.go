package main

import (
	"WS/system/db"
	"WS/system/route"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
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
