package users

import "net/http"

func registerRoute(path string, handler func(*AdminController)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		user := GetLoggedUser(r)
		if user == nil {
			http.NotFound(w, r)
			return
		}
		nc := NewAdminController(w, r, user)
		handler(nc)
	})
}

func RegisterRoutes() {
	registerRoute("/admin_users/all", (*AdminController).GetUsers)
	registerRoute("/admin_users/count_users", (*AdminController).CountUsers)
	registerRoute("/admin_users/admin_users", (*AdminController).CountAdminUsers)
	registerRoute("/admin_users/get_admin_users", (*AdminController).GetAdminUsers)
}
