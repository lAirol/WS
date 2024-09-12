package cookie

import (
	"net/http"
)

func CreateCookie(name, value string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return cookie
}

func GetIdCookie(r *http.Request) string {
	cookie, err := r.Cookie("Uid")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func GetUserAgentCookie(r *http.Request) string {
	//тут надо через r.header.Get но так падает приложение почему-то
	//cookie, err := r.Cookie("User-Agent")
	//if err != nil {
	//	return ""
	//}
	//return cookie.Value
	uAgent := r.Header.Get("User-Agent")
	return uAgent
}
