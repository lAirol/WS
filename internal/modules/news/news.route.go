package news

import (
	"WS/internal/modules/users"
	"net/http"
)

func registerRoute(path string, handler func(*AdminNewsController)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		nc := NewAdminNewsController(w, r, users.GetCurrUser(r))
		handler(nc)
	})
}

func RegisterRoutes() {
	registerRoute("/admin_news/config", (*AdminNewsController).NewsConfig)
	registerRoute("/admin_news/reply", (*AdminNewsController).NewsReply)
	registerRoute("/admin_news/identity", (*AdminNewsController).NewsIdentity)
	registerRoute("/admin_news/core", (*AdminNewsController).NewsCore)
}
