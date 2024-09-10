package users

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type IClient interface {
	GetType() bool
}

type Client struct {
	ID   string
	Conn *websocket.Conn
	Mu   sync.Mutex
	Type bool //false - юзер true - человечина с доступом к админке
}

type UserClient struct {
	*Client
}

func (u *UserClient) GetType() bool {
	return u.Type
}

type AdminClient struct {
	*Client
	someAdminFieldsLikeRules string
	UserAgent                string
}

func (a *AdminClient) GetType() bool {
	return a.Type
}

type Clients struct {
	UsersClients  map[*websocket.Conn]*UserClient
	AdminsClients map[string]*AdminClient
}

var CurrClients = Clients{
	UsersClients:  make(map[*websocket.Conn]*UserClient),
	AdminsClients: make(map[string]*AdminClient),
}

func GetLoggedUser(r *http.Request) (*AdminClient, error) {
	cookie, err := r.Cookie("Uid")
	Uid := cookie.Value
	if err != nil {
		return nil, err
	}
	return CurrClients.AdminsClients[Uid], nil
}
