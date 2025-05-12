package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"rt_forum/backend/models"
	"strings"

	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var users map[int]*websocket.Conn

func HandleWS(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	Conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"error": "can't upgrage the protocol",
		})
		return
	}
	token, err := r.Cookie("token")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"error": "can't upgrage the protocol",
		})
		return
	}
	id, err := models.GetId(db, strings.TrimPrefix(token.String(), "token="))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"error": "can't upgrage the protocol",
		})
		return
	}
	users[id]=Conn
	for _, val := range users{
		if val!= Conn {
			val.WriteJSON(map[string]interface{}{
				"id": id,
			})
		}
	}

	
}
