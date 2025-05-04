package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"rt_forum/backend/objects"

	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var User objects.User
func HandleWS(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	Conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"error": "can't upgrage the protocol",
		})
	}
	User.Connection = Conn
}
