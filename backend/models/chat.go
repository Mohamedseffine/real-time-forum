package models

import (
	"database/sql"
	"html"
	"log"
	"time"

	"rt_forum/backend/objects"
)

func GetChat(db *sql.DB, senderId int, recieverId int, lastInsertedId int) (objects.Chat, error) {
	var chat objects.Chat

	stm, err := db.Prepare(`SELECT * FROM messages WHERE ( ( sender_id = ? AND receiver_id = ? ) OR ( sender_id = ? AND receiver_id = ? ) ) AND id < ?  ORDER BY recieved_at DESC LIMIT 10`)
	if err != nil {
		log.Println("1", err)
		return objects.Chat{}, err
	}
	rows, err := stm.Query(senderId, recieverId, recieverId, senderId, lastInsertedId)
	if err != nil {
		log.Println("2", err)
		return objects.Chat{}, err
	}
	for rows.Next() {
		var message objects.Message
		err = rows.Scan(&message.MessageId, &message.UserId, &message.RecieverId, &message.Content, &message.Type, &message.Date, &message.Username, &message.Reciever)
		if err != nil {
			log.Println("3", err)
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
	res, err := stm.Exec(Data.UserId, Data.RecieverId, html.EscapeString(Data.Message), Data.Status, time.Now(), html.EscapeString(Data.Username), html.EscapeString(Data.Reciever_username))
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func GetBaseChat(db *sql.DB, senderId int, recieverId int) (objects.Chat, error) {
	var chat objects.Chat

	stm, err := db.Prepare(`SELECT * FROM messages WHERE  ( sender_id = ? AND receiver_id = ? ) OR ( sender_id = ? AND receiver_id = ? )  ORDER BY recieved_at DESC LIMIT 10`)
	if err != nil {
		log.Println("1", err)
		return objects.Chat{}, err
	}
	rows, err := stm.Query(senderId, recieverId, recieverId, senderId)
	if err != nil {
		log.Println("2", err)
		return objects.Chat{}, err
	}
	for rows.Next() {
		var message objects.Message
		err = rows.Scan(&message.MessageId, &message.UserId, &message.RecieverId, &message.Content, &message.Type, &message.Date, &message.Username, &message.Reciever)
		if err != nil {
			log.Println("3", err)
			return objects.Chat{}, err
		}
		message.Content = html.UnescapeString(message.Content)
		chat.Messages = append(chat.Messages, message)
	}
	return chat, nil
}



func UnreadMess(db *sql.DB, id int) ([]int, error) {
	stm, err := db.Prepare(`SELECT DISTINCT sender_id FROM messages WHERE receiver_id = ? AND mtype = "unread"`)

	if err != nil {
		return nil, err
	}
	ids := []int{}
	rows, err := stm.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var Uid int
		err := rows.Scan(&Uid)
		if err != nil {
			return nil, err
		}
		ids = append(ids, Uid)
	}
	return ids, nil
}