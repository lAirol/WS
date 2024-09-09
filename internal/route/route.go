package route

import (
	"WS/internal/handlers"
	"WS/internal/modules/admin"
	"WS/internal/modules/news"
	"WS/internal/modules/users"
	"net/http"
	"strings"
)

func InitRoute() {
	defRoutes()
	news.RegisterRoutes()
	admin.RegisterRoutes()
	users.RegisterRoutes()
	reqType()
}

func defRoutes() {
	http.HandleFunc("/ws", handlers.HandleConnections)
	http.HandleFunc("/wsadmin", handlers.HandleAdminConnections)
	http.HandleFunc("/index", handlers.HandleIndex)
}

func reqType() {
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
