package route

import (
	"WS/system/handlers"
	"net/http"
	"strings"
)

func InitRoute() {
	defRoutes()
	reqType()
}

func defRoutes() {
	http.HandleFunc("/ws", handlers.HandleConnections)
	http.HandleFunc("/index", handlers.HandleIndex)
}

func reqType() {

	// Обработка статических файлов
	staticHandler := http.FileServer(http.Dir("./public/"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if isStaticFileRequest(r) {
			// Запрос на статический файл
			cacheControlMiddleware(nil, staticHandler).ServeHTTP(w, r)
		} else {
			// Запрос на динамическую обработку
			handlers.HandleDynamicRoutes(w, r)
		}
	})
}

func isStaticFileRequest(r *http.Request) bool {
	return strings.Contains(r.URL.Path, ".")
}

func cacheControlMiddleware(next http.Handler, staticHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
