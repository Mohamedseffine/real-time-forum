package models

import (
	"database/sql"
	"html"
	"time"
)


func InsertComments(db *sql.DB, userid int, postid int, content string) (int, error) {
	stm, err := db.Prepare(`INSERT INTO comments (creator_id, post_id, creation_date, content) VALUES (?,?,?,?)`)
	if err != nil {
		return -1, err
	}	
	res, err := stm.Exec(userid, postid, time.Now(), html.EscapeString(content))
	if err != nil {
		return -1, err
	}
	id, err:=res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}