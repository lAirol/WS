package admin

import (
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

func (ac *AdminController) Load() {
	tmpl, err := template.ParseFiles("view.html")
	if err != nil {
		http.Error(ac.Writer, "Unable to load template", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title": "Admin Page",
		// Add more data here as needed
	}

	err = tmpl.Execute(ac.Writer, data)
	if err != nil {
		http.Error(ac.Writer, "Unable to render template", http.StatusInternalServerError)
	}
}
