package models

import (
	"database/sql"
	"rt_forum/backend/objects"
)


func InsertPost(db *sql.DB, postdata objects.Post) (int, error)  {
	stm , err := db.Prepare(`INSERT INTO posts (creator_id, title, creation_time, content) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return -1, err
	}
	res, err := stm.Exec(postdata.UserId, postdata.Title, postdata.Time, postdata.Content)
	if err != nil {
		return -1, err
	}
	id, err:=res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}



// func SelectCreatedPosts(db *sql.DB, )  {
// 	stm, err:=
// }