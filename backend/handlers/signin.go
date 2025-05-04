package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func Handlesignin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "method not allowed",
		})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&userdata); err != nil {
		log.Println("Failed to decode JSON:", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	loginpdatabase()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userdata)
}

func signinpdatabase() {
	whomi := userdata

	log.Println("login email :", whomi.Email)
	log.Println("login username :", whomi.Username)
	log.Println("login password :", whomi.Password)
}
