package news

import (
	"WS/internal/modules"
	"WS/internal/modules/users"
	"fmt"
	"net/http"
)

type AdminNewsController struct {
	modules.AdminController
}

func NewAdminNewsController(w http.ResponseWriter, r *http.Request, admin *users.AdminClient) *AdminNewsController {
	return &AdminNewsController{
		modules.AdminController{
			Request: r,
			Writer:  w,
			Admin:   admin,
		},
	}
}

func (nc *AdminNewsController) GetNews() {
	fmt.Print("1")
	// Логика обработки запроса на получение списка новостей
	// ...
}

func (nc *AdminNewsController) NewsConfig() {
	fmt.Print("2")
	// Логика обработки запроса на получение конфигурации новостей
	// ...
}

func (nc *AdminNewsController) NewsReply() {
	fmt.Print("3")
	// Логика обработки запроса на получение ответа на новость
	// ...
}

func (nc *AdminNewsController) NewsIdentity() {
	fmt.Print("4")
	// Логика обработки запроса на получение идентификации новости
	// ...
}

func (nc *AdminNewsController) NewsCore() {
	fmt.Print("5")
	// Логика обработки запроса на получение основной информации о новости
	// ...
}
