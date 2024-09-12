package users

import (
	"WS/internal/modules/response"
	"net/http"
)

type AdminController struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Admin   AdminClient
}

func NewAdminController(w http.ResponseWriter, r *http.Request, admin *AdminClient) *AdminController {
	return &AdminController{
		Request: r,
		Writer:  w,
		Admin:   *admin,
	}
}

func (ac *AdminController) CreateJSONResponse(data interface{}) {
	response.NewJSONResponse(ac.Writer, response.CreateResponse(data))
}

func (ac *AdminController) GetUsers() {
	ac.CreateJSONResponse(map[string]interface{}{"clients": len(CurrClients.UsersClients) + len(CurrClients.AdminsClients)})
}

func (ac *AdminController) CountUsers() {
	ac.CreateJSONResponse(map[string]interface{}{"clients": len(CurrClients.UsersClients)})
}

func (ac *AdminController) CountAdminUsers() {
	ac.CreateJSONResponse(map[string]interface{}{"clients": len(CurrClients.AdminsClients)})
}

func (ac *AdminController) GetAdminUsers() {
	ac.CreateJSONResponse(CurrClients.AdminsClients)
}
