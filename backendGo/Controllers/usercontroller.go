package controllers

import (
	driver "backendgo/Driver"
	structmodels "backendgo/StructModels"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user structmodels.NewUser

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error decode body Create User: %s /n", err)
		http.Error(w, "Invalid payload", 400)
	}

	if driver.UserExists(user.Username) {
		log.Printf("Error Creating User, Username already exists:  /n")
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	_, err := driver.CreateUser(user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("Error Creating User in driver: %s /n", err)
		http.Error(w, "Couldnt create user", 404)
		return
	}

}
