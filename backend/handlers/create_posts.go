package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"rt_forum/backend/objects"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB)  {
	var post objects.Post
	err:=json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"invalid data",
		})
		return
	}
	
}