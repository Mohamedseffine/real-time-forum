package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"rt_forum/backend/models"
)

type data struct {
	PostId int `json:"postid"`
}

func RetrieveComments(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	// Get postid from URL query
	postIDStr := r.URL.Query().Get("postid")
	if postIDStr == "" {
		http.Error(w, "Missing postid parameter", http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid postid parameter", http.StatusBadRequest)
		return
	}

	comments, err := models.GetComments(db, postID)
	if err != nil {
		http.Error(w, "Failed to retrieve comments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
