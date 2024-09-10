package handlers

import (
	"WS/internal/db"
	"WS/internal/db/tables"
	"WS/internal/extentions/random"
	"WS/internal/modules/admin"
	"WS/internal/modules/login"
	"WS/internal/modules/users"
	"fmt"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/html/index.html")
}

func HandleAdminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		login.Load(w, r)
	} else {
		username := r.PostFormValue("username")
		hash, err := db.GetHashByUser(username)
		if err != nil {

		}
		res := login.CheckPasswordHash(r.PostFormValue("password"), hash)
		if res == false {

		} else {
			cId := random.GenerateUUID()
			adminClient := &users.AdminClient{Client: &users.Client{ID: cId, Conn: nil, Type: true}, UserAgent: r.Header.Get("User-Agent")}
			users.CurrClients.AdminsClients[cId] = adminClient
			ac := admin.NewAdminController(w, r)
			ac.Load(adminClient)
		}
	}
}

func HandleDynamicRoutes(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	if url == "/" {
		http.Redirect(w, r, "/index", http.StatusMovedPermanently)
		return
	}

	// Проверка наличия URL в базе данных
	str, err := db.FindUrlExists(url)
	if err != nil {
		// Обработка ошибки при проверке в базе данных
		http.Error(w, "Произошла ошибка в HandleDynamicRoutes", http.StatusInternalServerError)
		return
	}

	if str != nil {
		handleController(w, r, str)
	} else {
		// Запись не найдена
		http.NotFound(w, r)
	}
}

func handleController(w http.ResponseWriter, r *http.Request, s *tables.Structure) {
	controllerName, err := db.LoadModuleInfo(s.ID)
	if err != nil {
		// Обработка ошибки
		return
	}
	fmt.Println(controllerName)
}
