package handlers

import (
	"WS/internal/config/constants"
	"WS/internal/extentions/random"
	"WS/internal/modules/users"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

type Sender struct {
	users.IClient
	Msg    []byte
	Server bool
}

var (
	upgrader  websocket.Upgrader
	Broadcast = make(chan Sender)
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
		sender := <-Broadcast
		mu.Lock()
		if sender.Server == true || !sender.IClient.GetType() {
			for _, userClient := range users.CurrClients.UsersClients {
				userClient.Mu.Lock()
				if userClient.Conn != nil && userClient != sender.IClient {
					err := userClient.Conn.WriteMessage(websocket.TextMessage, sender.Msg)
					if err != nil {
						log.Printf("error sending message: %v", err)
						userClient.Conn.Close()
						delete(users.CurrClients.UsersClients, userClient.Conn)
					}
				}
				userClient.Mu.Unlock()
			}
		} else if sender.Server == true || sender.IClient.GetType() {
			for _, adminClient := range users.CurrClients.AdminsClients {
				adminClient.Mu.Lock()
				if adminClient.Conn != nil && adminClient != sender.IClient {
					err := adminClient.Conn.WriteMessage(websocket.TextMessage, sender.Msg)
					if err != nil {
						log.Printf("error sending message: %v", err)
						adminClient.Conn.Close()
						delete(users.CurrClients.AdminsClients, adminClient.ID)
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

	defaultClient := &users.UserClient{Client: &users.Client{ID: random.GenerateUUID(), Conn: ws, Type: false, Ip: r.RemoteAddr}}
	mu.Lock()
	users.CurrClients.UsersClients[ws] = defaultClient
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
					log.Printf("index page connection closed: %v", err)
				}
				break
			}
			s := Sender{
				IClient: defaultClient,
				Msg:     msg,
			}
			Broadcast <- s
		}
	}()
}

func HandleAdminConnections(w http.ResponseWriter, r *http.Request) {
	// тут должны быть проверка на то, залогинен пользователь или нет
	admin := users.GetCurrUser(r)
	if admin == nil {
		return
	} else {
		if admin.Timer != nil {
			admin.Timer.Stop()
			admin.Timer = nil
		}
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading connection: %v", err)
		return
	}

	admin.Mu = mu
	admin.Conn = ws

	go func() {
		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					log.Printf("unexpected close error: %v", err)
				} else {
					log.Printf("connection closed: %v", err)
					crateCloseChannel(admin)
				}
				break
			}
			s := Sender{
				IClient: users.CurrClients.AdminsClients[admin.ID],
				Msg:     msg,
			}
			Broadcast <- s
		}
	}()
}

func crateCloseChannel(admin *users.AdminClient) {
	admin.Timer = time.AfterFunc(constants.Constants.AdminTimeoutSeconds*time.Second, func() {
		log.Printf("Удаление Юзера")
		if admin.Conn != nil {
			admin.Conn.Close()
		}
		delete(users.CurrClients.AdminsClients, admin.ID)
	})
}
