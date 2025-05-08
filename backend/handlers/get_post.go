package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Post struct {
	ID           int    `json:"id"`
	CreatorID    int    `json:"creator_id"`
	Username     string `json:"username"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CreationTime string `json:"creation_time"`
}

func RetrievePosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "method not allowed",
		})
		return
	}

	rows, err := db.Query(`
		SELECT id, creator_id, username, title, content, creation_time
		FROM posts
		ORDER BY creation_time DESC
	`)
	if err != nil {
		http.Error(w, "Failed to query posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.CreatorID, &post.Username, &post.Title, &post.Content, &post.CreationTime); err != nil {
			http.Error(w, "Error scanning post", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
