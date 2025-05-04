package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"rt_forum/backend/objects"
)

func  HandleSignUp(w http.ResponseWriter, r *http.Request, db *sql.DB)  {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":"method not allowed",
		})
		return
	}
	var userdata objects.LogData
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(userdata)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userdata)
	if err != nil {
		log.Fatal(err)
	}
}

