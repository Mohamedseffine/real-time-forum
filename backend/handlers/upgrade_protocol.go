package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"rt_forum/backend/models"
	"rt_forum/backend/objects"
	"strings"
	"sync"

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

	defer Conn.Close()
	token, err := r.Cookie("token")
	if err != nil {
		Conn.WriteJSON(map[string]any{
			"error": "hbbhhhbbbhbhhbhbhbh",
		})
		return
	}
	id, err := models.GetId(db, strings.TrimPrefix(token.String(), "token="))

	if err != nil {
		Conn.WriteJSON(map[string]any{
			"error": "pppppppppppppppppppppp",
		})
		return
	}
	mu.Lock()
	objects.Users[id] = Conn
	mu.Unlock()
	mu.Lock()
	defer delete(objects.Users, id)
	mu.Unlock()
	if len(objects.Users) > 1 {
		for _, val := range objects.Users {
			if val != Conn {
				val.WriteJSON(map[string]any{
					"id": id,
				})
			}
		}
	}
	var data objects.WsData
	users, err := models.GetAllUsers(db)
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
	for {
		var message objects.WsData
		err := Conn.ReadJSON(&message)
		if err != nil {
			Conn.WriteJSON(map[string]any{
				"error": "can not read message",
			})
			return
		}
		if len(objects.Users) > 1 {
			for _, val := range objects.Users {
				if val != Conn {
					val.WriteJSON(map[string]any{
						"message": message,
					})
				}
			}
		}
	}

}
