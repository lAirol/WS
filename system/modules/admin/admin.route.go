package admin

import "net/http"

func RegisterRoutes() {
	routes := map[string]func(*AdminController){
		"/admin": (*AdminController).Load,
	}

	for route, handler := range routes {
		http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			nc := NewAdminController(w, r)
			handler(nc)
		})
	}
}
