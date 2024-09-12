package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func CreateResponse(data interface{}) *Response {
	return &Response{
		Status:  200,
		Message: "ok",
		Data:    data,
	}
}

func CreateErrResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Write([]byte(`{"status": "` + string(rune(status)) + `", "message": "` + err.Error() + `"}`))
}

func NewJSONResponse(w http.ResponseWriter, response *Response) {
	jsonData, err := json.Marshal(response)
	if err != nil {
		CreateErrResponse(w, http.StatusInternalServerError, err)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
