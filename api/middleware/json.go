package middleware

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter) {
	OkResponse(w, 400, map[string]string{"error": "An error occured"})
}

func OkResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}