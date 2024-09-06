package handlers

import (
	"WS/system/modules/admin"
	"WS/system/modules/news"
	"net/http"
	"strings"
)

func InitRoute() {
	defRoutes()
	reqType()
	news.RegisterRoutes()
	admin.RegisterRoutes()
}

func defRoutes() {
	http.HandleFunc("/ws", HandleConnections)
	http.HandleFunc("/index", HandleIndex)
	http.HandleFunc("/admin", HandleAdmin)
}

func reqType() {
	staticHandler := http.FileServer(http.Dir("./public/"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if isStaticFileRequest(r) {
			// Запрос на статический файл
			cacheControlMiddleware(nil, staticHandler).ServeHTTP(w, r)
		} else if strings.HasPrefix(r.URL.Path, "/admin_") {
			// Запрос на путь, начинающийся с /admin_
			HandleAdminRoutes(w, r)
		} else {
			// Запрос на динамическую обработку
			HandleDynamicRoutes(w, r)
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
