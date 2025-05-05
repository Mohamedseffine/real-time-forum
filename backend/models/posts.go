package models

import (
	"database/sql"
	"time"

	"rt_forum/backend/objects"
)

func InsertPost(db *sql.DB, postdata objects.Post) (int, error) {
	stm, err := db.Prepare(`INSERT INTO posts (creator_id, title, username, creation_time, content) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return -1, err
	}
	res, err := stm.Exec(postdata.UserId, postdata.Title, postdata.Username, time.Now(), postdata.Content)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

// func SelectCreatedPosts(db *sql.DB, )  {
//     stm, err:=
// }

func InsertCategories(db *sql.DB, postid int, categoryid int) error {
	stm, err := db.Prepare(`INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	_, err = stm.Exec(postid, categoryid)
	return err
}
