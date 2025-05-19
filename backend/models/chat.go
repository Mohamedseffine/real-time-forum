package models

import (
	"database/sql"
	"rt_forum/backend/objects"
	"time"
)

func GetChat(db *sql.DB, senderId int, recieverId int, lastInsertedId int) (objects.Chat, error) {
	var chat objects.Chat
	stm, err := db.Prepare(`SELECT * FROM messages WHERE sender_id = ? AND receiver_id = ? AND id < ?  ORDER BY creation_date DESC LIMIT 10`)
	if err != nil {
		return objects.Chat{}, err
	}
	rows, err := stm.Query(senderId, recieverId, lastInsertedId)
	if err != nil {
		return objects.Chat{}, err
	}
	for rows.Next() {
		var message objects.Message
		err = rows.Scan(&message.MessageId, &message.UserId, &message.RecieverId, &message.Content, &message.Type, &message.Date, &message.Username, &message.Reciever)
		if err != nil {
			return objects.Chat{}, err
		}
		chat.Messages = append(chat.Messages, message)
	}
	return chat, nil
}

func InsertMessage(db *sql.DB, Data objects.WsData) (int, error) {
	stm, err := db.Prepare(`INSERT INTO messages (sender_id, receiver_id, content, mtype, recieved_at, sender_username, reciever_username) VALUES (?,?,?,?,?,?,?)`)
	if err != nil {
		return -1, err
	}
	res, err := stm.Exec(Data.UserId, Data.RecieverId, Data.Message, "", time.Now(), Data.Username, Data.Reciever_username)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}
