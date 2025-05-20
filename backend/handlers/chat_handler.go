package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"rt_forum/backend/models"
	"rt_forum/backend/objects"
)

type data struct {
	Sender_id    int `json:"sender_id"`
	Reciever_id  int `json:"receiver_id"`
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

		return
	}
	var messages objects.Chat
	if data.LastInsertId != 0 {
		messages, err = models.GetChat(db, data.Sender_id, data.Reciever_id, data.LastInsertId)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "database error",
			})
			return
		}
	}else {
		messages, err = models.GetBaseChat(db, data.Sender_id, data.Reciever_id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"error": "database error",
			})
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(messages)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"error": "sending data error",
		})
		return
	}
}
