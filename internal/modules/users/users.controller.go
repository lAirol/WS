package users

import (
	"fmt"
	"net/http"
)

type UsersController struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

func NewUsersController(w http.ResponseWriter, r *http.Request) *UsersController {
	return &UsersController{
		Request: r,
		Writer:  w,
	}
}

func (*UsersController) GetUsers() {
	fmt.Print(CurrClients.AdminsClients)
}

func (*UsersController) CountUsers() {

}

func (*UsersController) GetAdminUsers() {

}

func (*UsersController) GetDefaultUsers() {

}
