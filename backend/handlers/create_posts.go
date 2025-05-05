package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"rt_forum/backend/objects"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB)  {
	var post objects.Post
	json.NewDecoder(r.Body).Decode(&post)
}