package users

import "net/http"

func registerRoute(path string, handler func(*AdminUsersController)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		user := GetCurrUser(r)
		if user == nil {
			http.NotFound(w, r)
			return
		}
		nc := NewAdminUsersController(w, r, user)
		handler(nc)
	})
}

func RegisterRoutes() {
	registerRoute("/admin_users/all", (*AdminUsersController).GetUsers)
	registerRoute("/admin_users/count_users", (*AdminUsersController).CountUsers)
	registerRoute("/admin_users/count_admin_users", (*AdminUsersController).CountAdminUsers)
	registerRoute("/admin_users/get_admin_users", (*AdminUsersController).GetAdminUsers)
}
