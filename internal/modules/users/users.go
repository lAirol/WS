package users

import (
	"github.com/gorilla/websocket"
	"sync"
)

type IClient interface {
	IsLoggedIn() bool
}

type Client struct {
	ID     int
	Conn   *websocket.Conn
	Mu     sync.Mutex
	Logged bool
}

type UserClient struct {
	*Client
}

func (u *UserClient) IsLoggedIn() bool {
	return u.Logged
}

type AdminClient struct {
	*Client
	someAdminFieldsLikeRules string
}

func (a *AdminClient) IsLoggedIn() bool {
	return a.Logged
}

type Clients struct {
	UsersClients  map[*websocket.Conn]*UserClient
	AdminsClients map[*websocket.Conn]*AdminClient
}

var CurrClients = Clients{
	UsersClients:  make(map[*websocket.Conn]*UserClient),
	AdminsClients: make(map[*websocket.Conn]*AdminClient),
}
