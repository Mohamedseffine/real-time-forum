package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
		log.Println("1", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "can't parse data",
		})
		return
	}
	if !helpers.IsValidEmail(userdata.Email) {
		fmt.Println("1", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "invalid email",
		})
		return
	}
	if len(userdata.Password) < 8 {
		fmt.Println("2", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "too short password",
		})
		return
	}
	if !helpers.IsValidUesrname(userdata.Username) {
		fmt.Println("3", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "invalid username",
		})
		return
	}
	n := models.CheckUsername(db, userdata.Username)
	if n != 0 {
		log.Println(n)
		fmt.Println("4", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "username already used",
		})
		return
	}
	m := models.CheckEmail(db, userdata.Email)
	if m != 0 {
		log.Println(m)
		fmt.Println("5", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "email already used",
		})
		return
	}

	if !helpers.IsvalidName(userdata.Name) {
		fmt.Println("6", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "invalid name",
		})
		return
	}
	if !helpers.IsvalidName(userdata.FamilyName) {
		fmt.Println("7", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "invalid family name",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userdata.Password), 10)
	if err != nil {
		log.Println("2", err)
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
		log.Println("3", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "can't insert data",
		})
		return
	}
	User.Id = id
	token, err := uuid.NewV4()
	if err != nil {
		log.Println("4", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "can't insert data",
		})
		return
	}
	err = models.CreateSession(db, id, token.String(), time.Now(), time.Now().Add(2*time.Hour))
	if err != nil {
		log.Println("5", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "can't insert data",
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
		"token":token.String(),
	})
}
