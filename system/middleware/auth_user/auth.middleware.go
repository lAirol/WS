package auth_user

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверка, залогинен ли пользователь
		if !isLoggedIn(r) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isLoggedIn(r *http.Request) bool {
	// Логика проверки аутентификации пользователя
	// Например, проверка наличия сессии или токена
	// тут так же добавить продление токена если тип делать какой-нибудь запрос, т.е. если время жизни токена 30 мин, то устанавливать время жизни на текущее + 30 сек
	return r.Header.Get("Authorization") != ""
}
