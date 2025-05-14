package models

import (
	"database/sql"
	"rt_forum/backend/objects"
)

func GetChat(db *sql.DB, ChatId int, lastInsertedId int) (objects.Chat, error) {
	var chat objects.Chat
	stm, err := db.Prepare(`SELECT * FROM messages WHERE chat_id = ? AND id > ? LIMIT 10`)
	if err != nil {
		return objects.Chat{}, err
	}
	rows, err := stm.Query(ChatId, lastInsertedId)
	if err != nil {
		return objects.Chat{}, err
	}
	for rows.Next() {
		var message objects.Message
		err = rows.Scan(&message.MessageId, &message.UserId, &message.RecieverId, &message.Content, &message.Type, &message.Username, &message.Reciever)
		if err != nil {
			return objects.Chat{}, err
		}
		chat.Messages=append(chat.Messages, message)
	}
	return chat, nil
}
