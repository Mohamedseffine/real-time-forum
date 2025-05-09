package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"rt_forum/backend/models"
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
	log.Println(comment)
	if err != nil {
		log.Println("1", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"can't parse json",
		})
		return 
	}
	log.Println(comment.Username)
	id, err := models.InsertComments(db, comment.UserId,comment.PostId,comment.Content, comment.Username)
	if err != nil {
		log.Println("2", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"can't insert data",
		})
		return
	}
	comment.CommentId=id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)
}