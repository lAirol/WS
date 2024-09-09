package users

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Client interface {
	IsLoggedIn() bool
}

type UserClient struct {
	conn   *websocket.Conn
	mu     sync.Mutex
	logged bool
}

func (u *UserClient) IsLoggedIn() bool {
	return u.logged
}

type AdminClient struct {
	conn            *websocket.Conn
	mu              sync.Mutex
	logged          bool
	someAdminFields string
}

func (a *AdminClient) IsLoggedIn() bool {
	return a.logged
}
