package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"rt_forum/backend/models"
)

type logout struct {
	Id    int    `json:"id"`
	Username string `json:"username"`
	Token string `json:"token"`
}

func LogoutHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodDelete {
		fmt.Println(http.MethodDelete)
		fmt.Println(r.Method)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "method not allowed",
		})
		return
	}
	var data logout
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "method not allowed",
		})
		return
	}
	id := models.LogoutCheck(db, data.Token)
	if id != data.Id {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "session and id are not compatible",
		})
		return
	}
}
