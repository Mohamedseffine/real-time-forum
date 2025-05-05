package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"rt_forum/backend/helpers"
	"rt_forum/backend/models"
	"rt_forum/backend/objects"

	"github.com/gofrs/uuid"
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
		log.Println("1",err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "can't parse data",
		})
		return
	}
	if !helpers.IsValidEmail(userdata.Email) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "invalid email",
		})
		return
	}
	if !helpers.IsValidUesrname(userdata.Username) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "invalid username",
		})
		return
	}
	if !helpers.IsvalidName(userdata.Name) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "invalid name",
		})
		return
	}
	if !helpers.IsvalidName(userdata.FamilyName) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "invalid family name",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(userdata.Password), 10)
	if err != nil {
		log.Println("2",err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "can't hash password",
		})
		return
	}
	userdata.Password = string(hash)
	id, err := models.InsertUser(db, userdata)
	if err != nil {
		log.Println("3",err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "can't insert data",
		})
		return
	}
	User.Id = int(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userdata)
	if err != nil {
		log.Fatal(err)
	}

	token, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
	}
	err = models.CreateSession(db, id, token.String(), time.Now(), time.Now().Add(2*time.Hour))
	if err != nil {
		log.Println(err)
	}
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: token.String(),
		Expires: time.Now().Add(2*time.Hour),
		HttpOnly: true,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":id,
		"username":userdata.Username,
	})
}
