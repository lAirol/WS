package login

import (
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type Session struct {
}

func Load(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./internal/views/admin/login.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, "")
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
