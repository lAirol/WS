package admin

import (
	"WS/internal/extentions/cookie"
	"WS/internal/modules"
	"WS/internal/modules/users"
	"html/template"
	"net/http"
)

type AdminPageController struct {
	modules.AdminController
}

func NewAdminController(w http.ResponseWriter, r *http.Request, admin *users.AdminClient) *AdminPageController {
	return &AdminPageController{
		modules.AdminController{
			Request: r,
			Writer:  w,
			Admin:   admin,
		},
	}
}

func (ac *AdminPageController) Load() {
	tmpl, err := template.ParseFiles("./internal/views/admin/qwe.html", "./internal/views/admin/chat.html")
	if err != nil {
		http.Error(ac.Writer, "Unable to load template", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		//"clientId": clientId,
	}

	http.SetCookie(ac.Writer, cookie.CreateCookie("Uid", ac.Admin.ID))

	err = tmpl.Execute(ac.Writer, data)
	if err != nil {
		http.Error(ac.Writer, "Unable to render template", http.StatusInternalServerError)
	}
}
