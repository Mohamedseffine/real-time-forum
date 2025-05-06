package models

import (
	"database/sql"
	"time"
)


func InsertComments(db *sql.DB, userid int, postid int, content string) error {
	stm, err := db.Prepare(`INSERT INTO comments (creator_id, post_id, creation_date, content) VALUES (?,?,?,?)`)
	if err != nil {
		return err
	}	
	_, err = stm.Exec(userid, postid, time.Now(), content)
	if err != nil {
		return err
	}
	
	return nil
}