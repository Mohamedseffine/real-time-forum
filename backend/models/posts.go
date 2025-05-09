package models

import (
	"database/sql"
	"html"
	"time"

	"rt_forum/backend/objects"
)

func InsertPost(db *sql.DB, postdata objects.Post) (int, error) {
	stm, err := db.Prepare(`INSERT INTO posts (creator_id, title, username, creation_time, content) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return -1, err
	}
	res, err := stm.Exec(postdata.UserId, html.EscapeString(postdata.Title), postdata.Username, time.Now(), html.EscapeString(postdata.Content))
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

func GetPosts(db *sql.DB) ([]objects.Post, error) {
	rows, err := db.Query(`
		SELECT id, creator_id, username, title, content, creation_time
		FROM posts
		ORDER BY creation_time DESC
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var posts []objects.Post
	for rows.Next() {
		var post objects.Post
		if err := rows.Scan(&post.ID, &post.UserId, &post.Username, &post.Title, &post.Content, &post.CreationTime); err != nil {
			return nil, err
		}
		row, err := db.Query(`SELECT category_id FROM post_categories WHERE post_id = ?`, post.ID)
		if err != nil {
			return nil, err
		}
		for row.Next() {
			var id int
			name := ""
			if err = row.Scan(&id); err != nil {
				return nil, err
			}

			stm, err := db.Prepare(`SELECT category FROM categories WHERE id = ?`)
			if err != nil {
				return nil, err
			}
			stm.QueryRow(id).Scan(&name)
			// log.Println(name)
			post.Categorie = post.Categorie + " | " + name

		}
		// log.Println(post.Categorie)
		posts = append(posts, post)
	}
	return posts, nil
}
