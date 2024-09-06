package handlers

import (
	"WS/system/db"
	"WS/system/db/tables"
	"fmt"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/html/index.html")
}

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}

func HandleAdminRoutes(w http.ResponseWriter, r *http.Request) {

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
