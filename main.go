package main

import (
	"WS/system/db"
	"WS/system/handlers"
	"fmt"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{}

func main() {
	// Инициализация базы данных
	db.InitDB()

	// Маршрутизация
	http.HandleFunc("/ws", handlers.HandleConnections)
	http.HandleFunc("/index", handlers.HandleIndex)

	// Обработка статических файлов
	staticHandler := http.FileServer(http.Dir("./public/"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if isStaticFileRequest(r) {
			// Запрос на статический файл
			cacheControlMiddleware(nil, staticHandler).ServeHTTP(w, r)
		} else {
			// Запрос на динамическую обработку
			handleDynamicRoutes(w, r)
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func cacheControlMiddleware(next http.Handler, staticHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/index", http.StatusMovedPermanently)
			return
		}

		// Устанавливаем заголовок Cache-Control для кэширования на 5 минут
		w.Header().Set("Cache-Control", "public, max-age=300")

		// Вызываем соответствующий обработчик
		if isStaticFileRequest(r) {
			staticHandler.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func handleDynamicRoutes(w http.ResponseWriter, r *http.Request) {
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
		//handlers.HandleDynamicPage(w, r, url)
		fmt.Println("123")
	} else {
		// Если URL не найден, то возвращаем 404
		http.NotFound(w, r)
	}
}

func isStaticFileRequest(r *http.Request) bool {
	return strings.Contains(r.URL.Path, ".")
}
