package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
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
