package users

import (
	"WS/internal/extentions/cookie"
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
	UsersClients  map[*websocket.Conn]*UserClient `json:"user_clients"`
	AdminsClients map[string]*AdminClient         `json:"admin_clients"`
}

var CurrClients = Clients{
	UsersClients:  make(map[*websocket.Conn]*UserClient),
	AdminsClients: make(map[string]*AdminClient),
}

func GetLoggedUser(r *http.Request) *AdminClient {
	Uid := cookie.GetIdCookie(r)
	uAgent := cookie.GetUserAgentCookie(r)
	if adminExists(Uid) && CurrClients.AdminsClients[Uid].UserAgent == uAgent || uAgent != "" {
		return CurrClients.AdminsClients[Uid]
	}
	return nil
}

func adminExists(Uid string) bool {
	_, ok := CurrClients.AdminsClients[Uid]
	return ok
}
