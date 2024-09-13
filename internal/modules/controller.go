package modules

import (
	"WS/internal/modules/users"
	"net/http"
)

type Controller struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

type AdminController struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Admin   *users.AdminClient
}
