package models

import (
	"database/sql"
	"fmt"
	"log"
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
	query := `SELECT id, password FROM users WHERE username = ?`
	if typ == "email" {
		query = `SELECT id, password FROM users WHERE email = ?`
	}else if typ!="username"{
		return -1, "invalid login type"
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


func CheckUsername(db *sql.DB, username string) int {
	stm, err := db.Prepare(`SELECT COUNT(*) FROM users WHERE username = ?`)
	if err != nil {
		return -1
	}
	var n int
	err = stm.QueryRow(username).Scan(&n)
	if err != nil {
		return -1
	}
	
	return int(n)
}

func CheckEmail(db *sql.DB, email string) int {
	stm, err := db.Prepare(`SELECT COUNT(*) FROM users WHERE email = ?`)
	if err != nil {
		return -1
	}
	var n int
	err = stm.QueryRow(email).Scan(&n)
	if err != nil {
		return -1
	}
	return int(n)
}


func CheckSession(db *sql.DB,token string) (int) {
	var n int64
	stm, err := db.Prepare(`SELECT COUNT(*) FROM sessions WHERE token = ?`)
	if err != nil {
		log.Println("1",err)
		return -1
	}
	err = stm.QueryRow(token).Scan(&n)
	if err != nil {
		log.Println("1",err)
		return -1
	}
	return int(n)
}


func LogoutCheck(db *sql.DB, session string) (int) {
	stm, err := db.Prepare(`SELECT user_id FROM sessions WHERE token = ?`)
	if err != nil {
		return -1
	}
	var id int
	err =stm.QueryRow(session).Scan(&id)
	if err != nil {
		return -1
	}
	return id
}