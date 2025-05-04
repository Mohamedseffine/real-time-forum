package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"rt_forum/backend/models"
	"rt_forum/backend/objects"

	"golang.org/x/crypto/bcrypt"
)

func HandleSignUp(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
	hash, err := bcrypt.GenerateFromPassword([]byte(userdata.Password), 10)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "can't hash password",
		})
		return
	}
	userdata.Password = string(hash)
	_, err = models.InsertUser(db, userdata)
	if err!=nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "can't insert data",
		})
		return
	}
}

// w.Header().Set("Content-Type", "application/json")
// w.WriteHeader(http.StatusOK)
// err = json.NewEncoder(w).Encode(userdata)
// if err != nil {
// 	log.Fatal(err)
// }
