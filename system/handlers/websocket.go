package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"sync"
)

type client struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

type Sender struct {
	*client
	msg []byte
}

var (
	upgrader  websocket.Upgrader
	clients   = make(map[*websocket.Conn]*client)
	broadcast = make(chan Sender)
)

func init() {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	go broadcastMessages()
}

func broadcastMessages() {
	for {
		sender := <-broadcast
		for _, client := range clients {
			client.mu.Lock()
			fmt.Println(client)
			fmt.Println(&sender.client)
			fmt.Println(sender.client)
			if client.conn != nil && client != sender.client {
				err := client.conn.WriteMessage(websocket.TextMessage, sender.msg)
				if err != nil {
					log.Printf("error sending message: %v", err)
					client.conn.Close()
					delete(clients, client.conn)
				}
			}
			client.mu.Unlock()
		}
	}
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &client{conn: ws}
	clients[ws] = client

	go func() {
		defer ws.Close()
		defer delete(clients, ws)
		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					log.Printf("error: %v", err)
				} else {
					log.Printf("closed: %v", err)
				}
				break
			}
			// Обработка полученного сообщения (например, отправка в broadcast)
			s := Sender{
				client: client,
				msg:    msg,
			}
			broadcast <- s
		}
	}()
}
