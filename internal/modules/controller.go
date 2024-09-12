package modules

import "net/http"

type Controller struct {
	Request *http.Request
	Writer  http.ResponseWriter
}
