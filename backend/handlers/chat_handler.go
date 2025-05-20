package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rt_forum/backend/models"
)

type data struct {
	Sender_id    int `json:"sender_id"`
	Reciever_id  int `json:"reciever_id"`
	LastInsertId int `json:"last_id"`
}

func GetChatMessages(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]any{
			"error": "method not allowed",
		})
		return
	}
	var data data

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"error": "can not read json data",
		})
		// fmt.Println("lwala")
		return
	}
	messages, err := models.GetChat(db, data.Sender_id, data.Reciever_id, data.LastInsertId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"error": "database error",
		})
				fmt.Println(err)

		return
	}
	fmt.Println(messages)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(messages)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"error": "sending data error",
		})
				// fmt.Println("talta")

		return
	}
}
