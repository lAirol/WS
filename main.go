package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/", handleIndex)

	http.Handle("/public/", http.StripPrefix("/public/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/"+r.URL.Path)
		w.Header().Set("Cache-Control", "max-age=300") // Кэширование на 5 мин
	})))

	fmt.Println("Сервер запущен на http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что запрос действительно на WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	// Здесь будет ваш основной цикл обработки сообщений
	for {
		// Читаем сообщение от клиента
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("Получено сообщение:", string(msg))

		// Отправляем ответ клиенту
		err = ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/html/index.html")
}
