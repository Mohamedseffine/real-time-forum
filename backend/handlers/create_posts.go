package handlers

import (
	"database/sql"
	"encoding/json"
	"log"

	"net/http"

	"rt_forum/backend/models"
	"rt_forum/backend/objects"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "method not allowed",
		})
		return
	}
	var post objects.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid data",
		})
		return
	}

	id, err := models.InsertPost(db, post)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid data",
		})
		return
	}
	post.ID=id
	for _, value := range post.Categories {
		models.InsertCategories(db, id, value)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
