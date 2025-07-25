package handlers

import (
	"database/sql"
	"encoding/json"
	"log"

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
		log.Println(err, "line 37")
		Conn.WriteJSON(map[string]any{
			"error": "cookie error",
		})
		return
	}
	id, err := models.GetId(db, strings.TrimPrefix(token.String(), "token="))
	defer handleConnClosure(Conn, id)
	if err != nil {
		log.Println(err, "line 47")
		Conn.WriteJSON(map[string]any{
			"error": "db error ",
		})
		return
	}

	mu.Lock()
	objects.Users[id] = append(objects.Users[id], Conn)
	mu.Unlock()
	var data objects.WsData
	users, err := models.GetAllUsersBymessDate(db, id)
	if err != nil {
		log.Println(err, "line 70")
		Conn.WriteJSON(map[string]any{
			"error": err,
		})
		return
	}
	log.Println(users)
	usr, err := models.GetAllUsers(db, id)
	if err != nil {
		log.Println(err, "line 70")
		Conn.WriteJSON(map[string]any{
			"error": err,
		})
		return
	}
	users = append(users, usr...)
	for i := range users {
		if objects.Users[users[i].Id] != nil || len(objects.Users[users[i].Id]) != 0 {
			users[i].IsActive = 1
		}
	}
	log.Println(users)
	unreadMess, err := models.UnreadMess(db, id)
	if err != nil {
		log.Println(err, "line 53")
		Conn.WriteJSON(map[string]string{
			"error": "can't get messages",
		})
		return
	}
	data.Type = "all_users"
	data.Message = "sent"
	data.Users = users
	data.Unread = unreadMess

	Conn.WriteJSON(data)
	updateLoginState(id, data.Users)
	log.Println("things just ain't the same for ganhsters", objects.Users)
	for {
		var message objects.WsData
		err := Conn.ReadJSON(&message)
		if err != nil {
			log.Println(err, "line 90")
			Conn.WriteJSON(map[string]any{
				"error": "can not read message",
			})
			return
		}
		if message.Type == "message" {
			idm, err := models.InsertMessage(db, message)
			log.Println(" la ana lwl")
			if err != nil {
				log.Println(err, "line 101")
				Conn.WriteJSON(map[string]any{
					"error": err.Error(),
				})
				return
			}
			log.Println("3lach", id)

			for _, c := range objects.Users[id] {
				log.Println(c)
				if c != Conn {
					c.WriteJSON(map[string]any{
						"type":            message.Type,
						"id":              idm,
						"sender_id":       message.UserId,
						"sender_username": message.Username,
						"content":         message.Message,
						"time":            time.Now(),
						"status":          "unread",
						"reciever":        message.RecieverId,
					})
				}
			}
			SendMessage(message, idm, Conn)
		} else if message.Type == "update" {
			log.Println("ha ana ")
			err := models.UpdateMessState(db, message.UserId, message.RecieverId)
			log.Println("lwl")
			if err != nil {
				log.Println(err.Error(), "line 123")
				Conn.WriteJSON(map[string]string{
					"error": err.Error(),
				})
				return
			}
		} else if message.Type == "typing" {
			typing(message.UserId, message.RecieverId)
		}
	}

}

func handleConnClosure(Conn *websocket.Conn, id int) {
	sl := []*websocket.Conn{}
	mu.Lock()
	for _, v := range objects.Users[id] {
		if v != Conn {
			sl = append(sl, v)
		}
	}
	objects.Users[id] = sl
	mu.Unlock()
	Conn.Close()
	for ind, val := range objects.Users {
		if ind != id && len(objects.Users[id]) == 0 {
			for _, v := range val {
				v.WriteJSON(map[string]any{
					"type": "Disconneted",
					"id":   id,
				})
			}
		}
	}
}

func updateLoginState(id int, users []objects.Infos) {
	for ind, val := range objects.Users {
		if ind != id {
			for _, v := range val {
				v.WriteJSON(map[string]any{
					"type":  "connected",
					"id":    id,
					"users": users,
				})
			}

		}
	}
}

func SendMessage(message objects.WsData, id int, conn *websocket.Conn) {
	Conns := objects.Users[message.RecieverId]

	if len(Conns) != 0 {
		for _, v := range Conns {
			log.Println(Conns)
			v.WriteJSON(map[string]any{
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
}

func typing(senderId int, receiverId int) {
	if len(objects.Users[receiverId]) == 0 {
		return
	}
	for _, v := range objects.Users[receiverId] {
		v.WriteJSON(map[string]any{
			"type":   "typing",
			"sender": senderId,
		})
	}
}
