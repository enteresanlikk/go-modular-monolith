package common_presentation

import (
	"encoding/json"
	"net/http"
)

func SetJsonResponseHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func SetResponseStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func SetJsonResponseBody(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func JsonResponse(w http.ResponseWriter, data interface{}) {
	SetJsonResponseHeader(w)
	SetJsonResponseBody(w, data)
}

func JsonResponseWithStatus(w http.ResponseWriter, status int, data interface{}) {
	SetJsonResponseHeader(w)
	SetResponseStatus(w, status)
	SetJsonResponseBody(w, data)
}
