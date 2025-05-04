package models

import (
	"database/sql"
	"rt_forum/backend/objects"
	"time"
)


func InsertUser(db *sql.DB, user objects.LogData)(int64, error ) {
	stm, err :=db.Prepare(`INSERT INTO users(username, first_name, last_name, gender, email, password, creation_date) VALUES (?,?,?,?,?,?,?)`)
	if err != nil {
		return -1,  err
	}
	defer stm.Close()

	res , err := stm.Exec(user.Username, user.Name, user.FamilyName, user.Gender, user.Email, user.Password, time.Now())
	if err != nil {
		return -1 , err
	}
	lastid, err :=res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return lastid, nil
}


func ExtractUser(db *sql.DB,password string, log string, typ string)(int, error)  {
	var query =`SELECT (id) FROM USERS WHERE username = ?`
	if typ == "email" {
		query = `SELECT (id) FROM USERS WHERE email = ?`
	}
	stm, err := db.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stm.Close()
	

	return 5, nil
}