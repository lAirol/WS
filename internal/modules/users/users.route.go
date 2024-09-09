package users

import "net/http"

func registerRoute(path string, handler func(*UsersController)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		nc := NewUsersController(w, r)
		handler(nc)
	})
}

func RegisterRoutes() {
	registerRoute("/users/all", (*UsersController).GetUsers)
	registerRoute("/users/count", (*UsersController).CountUsers)
	registerRoute("/users/admin_users", (*UsersController).GetAdminUsers)
	registerRoute("/users/users", (*UsersController).GetDefaultUsers)
}
