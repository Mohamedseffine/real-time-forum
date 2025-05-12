package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"rt_forum/backend/models"
	"rt_forum/backend/objects"
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

var Users = make(map[int]*websocket.Conn, 24)

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
	Users[id] = Conn

	log.Println(Users)
	defer delete(Users, id)
	if len(Users) > 1 {
		for _, val := range Users {
			if val != Conn {
				val.WriteJSON(map[string]any{
					"id": id,
				})
			}
		}
	}
	users, err := models.GetAllUsers(db)
	if err != nil {
		log.Println(err)
		Conn.WriteJSON(map[string]any{
			"error": err,
		})
		return
	}
	Conn.WriteJSON(users)
	for {
		var message objects.WsData
		err := Conn.ReadJSON(&message)
		if err != nil {
			Conn.WriteJSON(map[string]any{
				"error": "can not read message",
			})
			return
		}
		if len(Users) > 1 {
			for _, val := range Users {
				if val != Conn {
					val.WriteJSON(map[string]any{
						"message": message,
					})
				}
			}
		}
	}

}
