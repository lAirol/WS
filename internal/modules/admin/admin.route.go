package admin

import (
	"net/http"
)

func registerRoute(path string, handler func(*AdminController)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		nc := NewAdminController(w, r)
		handler(nc)
	})
}

func RegisterRoutes() {
	//registerRoute("/admin/tree", (*AdminController).Load)
	//registerRoute("/admin1", (*AdminController).Load)
}
