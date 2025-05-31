package handlers

import (
	"database/sql"
	"encoding/json"

	"net/http"
	"rt_forum/backend/models"
	"rt_forum/backend/objects"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var mu sync.RWMutex

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
		Conn.WriteJSON(map[string]any{
			"error": "cookie error",
		})
		return
	}
	id, err := models.GetId(db, strings.TrimPrefix(token.String(), "token="))
	defer handleConnClosure(Conn, id)
	if err != nil {
		Conn.WriteJSON(map[string]any{
			"error": "db error ",
		})
		return
	}
	mu.Lock()
	objects.Users[id] = Conn
	mu.Unlock()

	// if len(objects.Users) > 1 {
	// 	for _, val := range objects.Users {
	// 		if val != Conn {
	// 			Conn.WriteJSON(map[string]any{
	// 				"type": "connected",
	// 				"id":   id,
	// 			})
	// 		}
	// 	}
	// }
	var data objects.WsData
	users, err := models.GetAllUsers(db, id)
	if err != nil {
		Conn.WriteJSON(map[string]any{
			"error": err,
		})
		return
	}
	for i := range users {
		if objects.Users[users[i].Id] != nil {

			users[i].IsActive = 1
		}
	}
	data.Type = "all_users"
	data.Message = "sent"
	data.Users = users
	Conn.WriteJSON(data)
	updateLoginState(Conn, id, data.Users)
	for {
		var message objects.WsData
		err := Conn.ReadJSON(&message)
		if err != nil {
			Conn.WriteJSON(map[string]any{
				"error": "can not read message",
			})
			return
		}
		if message.Type == "message" {
			id, err := models.InsertMessage(db, message)
			if err != nil {
				Conn.WriteJSON(map[string]any{
					"error": err.Error(),
				})
			}
			SendMessage(message, id)
		}
	}

}

func handleConnClosure(Conn *websocket.Conn, id int) {
	mu.Lock()
	delete(objects.Users, id)
	mu.Unlock()
	Conn.Close()
	for _, val := range objects.Users {
		if val != Conn {
			val.WriteJSON(map[string]any{
				"type": "Disconneted",
				"id":   id,
			})
		}
	}
}

func updateLoginState(Conn *websocket.Conn, id int, users []objects.Infos) {
	for _, val := range objects.Users {
		if val != Conn {
			val.WriteJSON(map[string]any{
				"type":  "connected",
				"id":    id,
				"users": users,
			})
		}
	}
}

func SendMessage(message objects.WsData, id int) {
	if Conn := objects.Users[message.RecieverId]; Conn != nil {
		Conn.WriteJSON(map[string]any{
			"type":            message.Type,
			"id":              id,
			"sender_id":       message.UserId,
			"sender_username": message.Username,
			"content":         message.Message,
			"time":            time.Now(),
			"status":          "unread",
		})
	}
}
