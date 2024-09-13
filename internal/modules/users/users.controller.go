package users

import (
	"WS/internal/modules/response"
	"net/http"
)

type AdminUsersController struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Admin   AdminClient
}

func NewAdminUsersController(w http.ResponseWriter, r *http.Request, admin *AdminClient) *AdminUsersController {
	return &AdminUsersController{
		Request: r,
		Writer:  w,
		Admin:   *admin,
	}
}

func (ac *AdminUsersController) CreateJSONResponse(data interface{}) {
	response.NewJSONResponse(ac.Writer, response.CreateResponse(data))
}

func (ac *AdminUsersController) GetUsers() {
	ac.CreateJSONResponse(map[string]interface{}{"clients": len(CurrClients.UsersClients) + len(CurrClients.AdminsClients)})
}

func (ac *AdminUsersController) CountUsers() {
	ac.CreateJSONResponse(map[string]interface{}{"clients": len(CurrClients.UsersClients)})
}

func (ac *AdminUsersController) CountAdminUsers() {
	ac.CreateJSONResponse(map[string]interface{}{"clients": len(CurrClients.AdminsClients)})
}

func (ac *AdminUsersController) GetAdminUsers() {
	ac.CreateJSONResponse(CurrClients.AdminsClients)
}
