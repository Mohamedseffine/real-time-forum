package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"rt_forum/backend/objects"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB)  {
	var comment objects.Comment
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"method not allowed",
		})
		return
	}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"can't parse json",
		})
		return 
	}
	
}