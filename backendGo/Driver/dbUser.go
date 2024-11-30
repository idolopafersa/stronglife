package driver

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, email, password string) (int, error) {

	passwordh, err := hashPassword(password)
	if err != nil {
		log.Printf("Couldnt hash password: %s /n", password)
		return -1, err
	}

	query := "INSERT INTO Users (username, email,password_hash) VALUES (? , ? , ?)"
	resul, err := db.Exec(query, username, email, passwordh)

	if err != nil {
		log.Println(resul, err)
		return -1, err
	} else {
		return 0, nil
	}

}

func UserExists(username string) bool {
	var user string

	query := "SELECT username FROM Users WHERE username = ?"

	// Query the database for the user and scan the result into the user struct
	err := db.QueryRow(query, username).Scan(&user)

	if err != nil {
		if err == sql.ErrNoRows {

			return false
		}
		log.Println("Error fetching user: %v\n", err)
		return false
	}

	return true
}

func CorrectPassword(username, password string) bool {
	var aldb []byte
	query := ("SELECT password_hash FROM Users WHERE username = ?")

	db.QueryRow(query, username).Scan(&aldb)

	if err := bcrypt.CompareHashAndPassword(aldb, []byte(password)); err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func GetUserID(name string) (int, error) {
	var id int
	fmt.Print(name)
	query := `Select id FROM Users WHERE username=?`
	err := db.QueryRow(query, name).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			print(err)
			return 0, err
		}
		print(err)
		return 0, err
	}

	return id, nil
}
