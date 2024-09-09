package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type client interface {
	isLoggedIn() bool
}

type userClient struct {
	conn   *websocket.Conn
	mu     sync.Mutex
	logged bool
}

func (u *userClient) isLoggedIn() bool {
	return u.logged
}

type Clients struct {
	UsersClients  map[*websocket.Conn]*userClient
	AdminsClients map[*websocket.Conn]*adminClient
}

type adminClient struct {
	conn            *websocket.Conn
	mu              sync.Mutex
	logged          bool
	someAdminFields string
}

func (a *adminClient) isLoggedIn() bool {
	return a.logged
}

type Sender struct {
	client
	msg []byte
}

var (
	upgrader  websocket.Upgrader
	broadcast = make(chan Sender)
	mu        sync.Mutex
	clients   = Clients{
		UsersClients:  make(map[*websocket.Conn]*userClient),
		AdminsClients: make(map[*websocket.Conn]*adminClient),
	}
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
		if !sender.client.isLoggedIn() {
			for _, userClient := range clients.UsersClients {
				userClient.mu.Lock()
				if userClient.conn != nil && userClient != sender.client {
					err := userClient.conn.WriteMessage(websocket.TextMessage, sender.msg)
					if err != nil {
						log.Printf("error sending message: %v", err)
						userClient.conn.Close()
						delete(clients.UsersClients, userClient.conn)
					}
				}
				userClient.mu.Unlock()
			}
		} else {
			for _, adminClient := range clients.AdminsClients {
				adminClient.mu.Lock()
				if adminClient.conn != nil && adminClient != sender.client {
					err := adminClient.conn.WriteMessage(websocket.TextMessage, sender.msg)
					if err != nil {
						log.Printf("error sending message: %v", err)
						adminClient.conn.Close()
						delete(clients.AdminsClients, adminClient.conn)
					}
				}
				adminClient.mu.Unlock()
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

	client := &userClient{conn: ws, logged: false}
	mu.Lock()
	clients.UsersClients[ws] = client
	mu.Unlock()

	go func() {
		defer func() {
			ws.Close()
			mu.Lock()
			delete(clients.UsersClients, ws)
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
				client: client,
				msg:    msg,
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

	adminClient := &adminClient{conn: ws, logged: true}
	mu.Lock()
	clients.AdminsClients[ws] = adminClient
	mu.Unlock()

	go func() {
		defer func() {
			ws.Close()
			mu.Lock()
			delete(clients.AdminsClients, ws)
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
				client: adminClient,
				msg:    msg,
			}
			broadcast <- s
		}
	}()
}
