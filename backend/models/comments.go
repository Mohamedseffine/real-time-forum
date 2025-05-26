package models

import (
	"database/sql"
	"html"
	"time"

	"rt_forum/backend/objects"
)

func InsertComments(db *sql.DB, userid int, postid int, content string, username string) (int, error) {
	stm, err := db.Prepare(`INSERT INTO comments (creator_id, post_id, username, creation_date, content) VALUES (?,?,?,?,?)`)
	if err != nil {
		return -1, err
	}
	res, err := stm.Exec(userid, postid, html.EscapeString(username), time.Now(), html.EscapeString(content))
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func GetComments(db *sql.DB, postid int) ([]objects.Comment, error) {
	rows, err := db.Query(`
		SELECT id, creator_id, post_id, username, creation_date, content 
		FROM comments 
		WHERE post_id = ?`, postid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []objects.Comment
	for rows.Next() {
		var comment objects.Comment
		err := rows.Scan(&comment.CommentId, &comment.UserId, &comment.PostId, &comment.Username, &comment.Time, &comment.Content)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
