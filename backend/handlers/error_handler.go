package handlers

import (
	"encoding/json"
	"net/http"
	"rt_forum/backend/objects"
)

func ErrorHandler(err objects.Error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"content":err.ErrorMessage,
	})
}
