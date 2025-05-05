package models

import (
	"database/sql"
	"fmt"
	"time"

	"rt_forum/backend/objects"

	"golang.org/x/crypto/bcrypt"
)

func InsertUser(db *sql.DB, user objects.LogData) (int, error) {
	stm, err := db.Prepare(`INSERT INTO users(username, first_name, last_name, gender, email, password, creation_date) VALUES (?,?,?,?,?,?,?)`)
	if err != nil {
		return -1, err
	}
	defer stm.Close()

	res, err := stm.Exec(user.Username, user.Name, user.FamilyName, user.Gender, user.Email, user.Password, time.Now())
	if err != nil {
		return -1, err
	}
	lastid, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(lastid), nil
}

func ExtractUser(db *sql.DB, password string, log string, typ string) (int, string) {
	query := `SELECT id, password FROM USERS WHERE username = ?`
	if typ == "email" {
		query = `SELECT id, password FROM USERS WHERE email = ?`
	}
	stm, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return -1, "database error"
	}
	defer stm.Close()
	var id int
	var hashpassword string
	err = stm.QueryRow(log).Scan(&id, &hashpassword)
	if err != nil {
		return -1, "invalid username"
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	if err != nil {
		return -1, "invalid password"
	}

	return int(id), ""
}

func CreateSession(db *sql.DB, id int, token string, creationTime time.Time, expiration time.Time) error {
	stm, err := db.Prepare(`INSERT INTO sessions (token, created_at, expires_at, user_id) VALUES (?,?,?,?)`)
	if err != nil {
		return err
	}
	_, err = stm.Exec(token, creationTime, expiration, id)

	return err
}
