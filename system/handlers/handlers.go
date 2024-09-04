package handlers

import (
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/html/index.html")
}
