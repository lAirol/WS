package news

import "net/http"

func RegisterRoutes() {
	routes := map[string]func(*NewsController){
		"/news":          (*NewsController).GetNews,
		"/news/config":   (*NewsController).NewsConfig,
		"/news/reply":    (*NewsController).NewsReply,
		"/news/identity": (*NewsController).NewsIdentity,
		"/news/core":     (*NewsController).NewsCore,
	}

	for route, handler := range routes {
		http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			nc := NewNewsController(w, r)
			handler(nc)
		})
	}

	//для уролв которые надо проверять на то, залогинен ли пользователь
	// import "WS/system/middleware/auth_user"
	//for route, handler := range routes {
	//	http.Handle(route, auth_user.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		nc := NewNewsController(w, r)
	//		handler(nc)
	//	})))
	//}
}
