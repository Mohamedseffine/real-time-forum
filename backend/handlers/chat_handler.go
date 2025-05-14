package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetChatMessages(w http.ResponseWriter, r *http.Request, db *sql.DB){
	if r.Method!= http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]any{
			"error":"method not allowed",
		})
		return
	}
	
}