package login

import (
	"WS/internal/extentions/cookie"
	"WS/internal/extentions/random"
	"WS/internal/modules/admin"
	"WS/internal/modules/users"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type Session struct {
}

func Load(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	tmpl, err := template.ParseFiles("./internal/views/admin/login.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func LoadAdminPanel(w http.ResponseWriter, r *http.Request) {
	if users.GetLoggedUser(r) != nil {
		admin.NewAdminController(w, r).Load(users.GetLoggedUser(r))
	} else {
		cId := random.GenerateUUID()
		adminClient := &users.AdminClient{Client: &users.Client{ID: cId, Conn: nil, Type: true}, UserAgent: cookie.GetUserAgentCookie(r)}
		users.CurrClients.AdminsClients[cId] = adminClient
		ac := admin.NewAdminController(w, r)
		ac.Load(adminClient)
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
