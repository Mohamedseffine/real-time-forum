package models

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	_"github.com/mattn/go-sqlite3"
)

var database *sql.DB

func DatabaseExec()*sql.DB {
	err := error(nil)
	database, err = sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		fmt.Println(" failed to open database: ", err)
		return nil
	}

	// Read the schema SQL file
	schema, err := ioutil.ReadFile("./database/schema.sql")
	if err != nil {
		fmt.Println(" failed to read schema file: ", err)
		return nil 
	}

	// Execute the SQL commands in the schema file
	_, err = database.Exec(string(schema))
	if err != nil {
		fmt.Println(" failed to execute schema:", err)
		return nil 
	}
	fmt.Println("ok dazt mzian gng")
	return database
}

// func CloseDatabase() {
// 	if Database != nil {
// 		err := Database.Close()
// 		if err != nil {
// 			fmt.Println("Error closing database:", err)
// 		} else {
// 			fmt.Println("Database closed successfully.")
// 		}
// 	}
// }
