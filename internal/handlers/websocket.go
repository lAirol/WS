package handlers

import (
	"WS/internal/modules/users"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Sender struct {
	users.IClient
	msg []byte
}

var (
	upgrader  websocket.Upgrader
	broadcast = make(chan Sender)
	mu        sync.Mutex
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
		mu.Lock()
		if !sender.IClient.IsLoggedIn() {
			for _, userClient := range users.CurrClients.UsersClients {
				userClient.Mu.Lock()
				if userClient.Conn != nil && userClient != sender.IClient {
					err := userClient.Conn.WriteMessage(websocket.TextMessage, sender.msg)
					if err != nil {
						log.Printf("error sending message: %v", err)
						userClient.Conn.Close()
						delete(users.CurrClients.UsersClients, userClient.Conn)
					}
				}
				userClient.Mu.Unlock()
			}
		} else {
			for _, adminClient := range users.CurrClients.AdminsClients {
				adminClient.Mu.Lock()
				if adminClient.Conn != nil && adminClient != sender.IClient {
					err := adminClient.Conn.WriteMessage(websocket.TextMessage, sender.msg)
					if err != nil {
						log.Printf("error sending message: %v", err)
						adminClient.Conn.Close()
						delete(users.CurrClients.AdminsClients, adminClient.Conn)
					}
				}
				adminClient.Mu.Unlock()
			}
		}

		mu.Unlock()
	}
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading connection: %v", err)
		return
	}

	Client := &users.UserClient{Conn: ws, Logged: false}
	mu.Lock()
	users.CurrClients.UsersClients[ws] = Client
	mu.Unlock()

	go func() {
		defer func() {
			ws.Close()
			mu.Lock()
			delete(users.CurrClients.UsersClients, ws)
			mu.Unlock()
		}()
		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					log.Printf("unexpected close error: %v", err)
				} else {
					log.Printf("connection closed: %v", err)
				}
				break
			}
			s := Sender{
				IClient: Client,
				msg:     msg,
			}
			broadcast <- s
		}
	}()
}

func HandleAdminConnections(w http.ResponseWriter, r *http.Request) {
	// тут должны быть проверка на то, залогинен пользователь или нет
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading connection: %v", err)
		return
	}

	adminClient := &users.AdminClient{Conn: ws, Logged: true}
	mu.Lock()
	users.CurrClients.AdminsClients[ws] = adminClient
	mu.Unlock()

	go func() {
		defer func() {
			ws.Close()
			mu.Lock()
			delete(users.CurrClients.AdminsClients, ws)
			mu.Unlock()
		}()
		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					log.Printf("unexpected close error: %v", err)
				} else {
					log.Printf("connection closed: %v", err)
				}
				break
			}
			s := Sender{
				IClient: adminClient,
				msg:     msg,
			}
			broadcast <- s
		}
	}()
}
