package handlers

import (
	"WS/internal/db"
	"WS/internal/db/tables"
	"WS/internal/modules/login"
	"WS/internal/modules/users"
	"fmt"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./internal/views/index.html")
}

func HandleAdminLogin(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if r.Method == http.MethodGet {
		if users.GetCurrUser(r) != nil {
			login.LoadAdminPanel(w, r)
		} else {
			login.Load(w, r, data)
		}
	} else {
		username := r.PostFormValue("username")
		hash, err := db.GetHashByUser(username)

		if err != nil {
			data["username_not_found"] = "Пользователь не найден"
			//тут можно начинать блокировать по ip
			login.Load(w, r, data)
			return
		} else if !login.CheckPasswordHash(r.PostFormValue("password"), hash) {
			//тут можно начинать блокировать по юзеру
			data["wrong_password"] = "Неверно введён пароль"
			login.Load(w, r, data)
		} else {
			login.LoadAdminPanel(w, r)
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
