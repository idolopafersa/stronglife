package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func OpenDB() {

	dbdirection := fmt.Sprintf("%s:%s@tcp(localhost:3306)/StrongLife", os.Getenv("Database_username"), os.Getenv("Database_password"))
	var err error
	db, err = sql.Open("mysql", dbdirection)

	fmt.Printf("connecting to... %s\n", dbdirection)
	if err != nil {
		pp := fmt.Sprintf("Couldn't open database.... %s", err)
		panic(pp)
	}

	// Ping the database to check if the connection is successful
	if err := db.Ping(); err != nil {
		log.Println(err)
		panic("Not opened")
	}
	fmt.Println("Connected to MariaDB!")

}

func CloseDB() {
	if db != nil {
		db.Close()
	} else {
		log.Println("Couldnt close database")
	}
}
