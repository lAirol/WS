package news

import "net/http"

func registerRoute(path string, handler func(*NewsController)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		nc := NewNewsController(w, r)
		handler(nc)
	})
}

func RegisterRoutes() {
	registerRoute("/news/config", (*NewsController).NewsConfig)
	registerRoute("/news/reply", (*NewsController).NewsReply)
	registerRoute("/news/identity", (*NewsController).NewsIdentity)
	registerRoute("/news/core", (*NewsController).NewsCore)
}
