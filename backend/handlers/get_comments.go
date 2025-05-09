package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"rt_forum/backend/models"
)
type data struct{
	PostId int `json:"postid"`
}
 
func RetrieveComments(w http.ResponseWriter, r *http.Request, db *sql.DB)  {
	if r.Method!=http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"invalid method",
		})
		return
	}
	var data data
	json.NewDecoder(r.Body).Decode(&data)
	comments, err := models.GetComments(db, data.PostId)
	if err!=nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"invalid method",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}