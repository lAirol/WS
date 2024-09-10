package admin

import (
	"WS/internal/extentions/cookie"
	"WS/internal/modules/users"
	"html/template"
	"net/http"
)

type AdminController struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

func NewAdminController(w http.ResponseWriter, r *http.Request) *AdminController {
	return &AdminController{
		Request: r,
		Writer:  w,
	}
}

func (ac *AdminController) Load(adminClient *users.AdminClient) {
	tmpl, err := template.ParseFiles("./internal/views/admin/qwe.html", "./public/html/chat.html")
	if err != nil {
		http.Error(ac.Writer, "Unable to load template", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		//"clientId": clientId,
	}

	http.SetCookie(ac.Writer, cookie.CreateCookie("Uid", adminClient.ID))

	err = tmpl.Execute(ac.Writer, data)
	if err != nil {
		http.Error(ac.Writer, "Unable to render template", http.StatusInternalServerError)
	}
}
