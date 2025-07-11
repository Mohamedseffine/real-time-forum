package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sync"

	"rt_forum/backend/models"
	"rt_forum/backend/objects"
)

type logout struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func LogoutHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodDelete {
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "can't parse data",
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
	var mu = sync.Mutex{}
	mu.Lock()
	delete(objects.Users, id)
	mu.Unlock()
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	})
	err = models.DeleteSession(db, data.Token)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "server error",
		})
		return
	}
}
