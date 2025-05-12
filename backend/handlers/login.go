package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"rt_forum/backend/models"
	"rt_forum/backend/objects"

	"github.com/gofrs/uuid"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "method not allowed",
		})
		return
	}

	var userdata objects.LogData
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "can't parse data",
		})
		return
	}

	id, erro := models.ExtractUser(db, userdata.Password, userdata.Username, userdata.LogType)
	if erro != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": erro,
		})
		return
	}
	token, err := uuid.NewV4()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "internal server error",
		})
		return
	}
	err = models.CreateSession(db, id, token.String(), time.Now(), time.Now().Add(2*time.Hour))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": string(err.Error()),
		})
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token.String(),
		Expires:  time.Now().Add(2 * time.Hour),
		HttpOnly: true,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":       id,
		"username": userdata.Username,
		"token":    token.String(),
	})
}
