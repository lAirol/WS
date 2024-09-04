package handlers

import (
	"WS/system/db"
	"fmt"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/html/index.html")
}

func HandleDynamicPage(w http.ResponseWriter, r *http.Request, url string) {

}

func HandleDynamicRoutes(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index", http.StatusMovedPermanently)
		return
	}
	// Проверка наличия URL в базе данных
	url := r.URL.Path
	exists, err := db.CheckUrlExists(url)
	if err != nil {
		// Обработка ошибки при проверке в базе данных
		http.Error(w, "Произошла ошибка", http.StatusInternalServerError)
		return
	}

	if exists {
		// Если URL найден в базе, то обрабатываем его
		// Например, вызываем соответствующую функцию для генерации страницы
		HandleDynamicPage(w, r, url)
		fmt.Println("123")
	} else {
		// Если URL не найден, то возвращаем 404
		http.NotFound(w, r)
	}
}
