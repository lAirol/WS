package news

import (
	"fmt"
	"net/http"
)

type NewsController struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

func NewNewsController(w http.ResponseWriter, r *http.Request) *NewsController {
	return &NewsController{
		Request: r,
		Writer:  w,
	}
}

func (nc *NewsController) GetNews() {
	fmt.Print("1")
	// Логика обработки запроса на получение списка новостей
	// ...
}

func (nc *NewsController) NewsConfig() {
	fmt.Print("2")
	// Логика обработки запроса на получение конфигурации новостей
	// ...
}

func (nc *NewsController) NewsReply() {
	fmt.Print("3")
	// Логика обработки запроса на получение ответа на новость
	// ...
}

func (nc *NewsController) NewsIdentity() {
	fmt.Print("4")
	// Логика обработки запроса на получение идентификации новости
	// ...
}

func (nc *NewsController) NewsCore() {
	fmt.Print("5")
	// Логика обработки запроса на получение основной информации о новости
	// ...
}
