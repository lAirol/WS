package news

import "net/http"

func registerRoute(path string, handler func(*NewsController)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		nc := NewNewsController(w, r)
		handler(nc)
	})
}

func RegisterRoutes() {
	registerRoute("/admin_news/config", (*NewsController).NewsConfig)
	registerRoute("/admin_news/reply", (*NewsController).NewsReply)
	registerRoute("/admin_news/identity", (*NewsController).NewsIdentity)
	registerRoute("/admin_news/core", (*NewsController).NewsCore)
}
