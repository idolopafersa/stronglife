package controllers

import (
	driver "backendgo/Driver"
	structmodels "backendgo/StructModels"
	"backendgo/security"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
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

func Login(w http.ResponseWriter, r *http.Request) {
	var creds structmodels.Credentials
	var id int
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.Printf("Error Decoding Body in Login function: %s /n", err)
		http.Error(w, "Error Decoding Body", http.StatusBadRequest)
		return
	}

	if driver.UserExists(creds.Username) && driver.CorrectPassword(creds.Username, creds.Password) {
		id, _ = driver.GetUserID(creds.Username)
		token, _ := security.CreateToken(id)
		//when a user sign in, they will receive a cookie to keep their JWT
		cookie := http.Cookie{
			Name:     "jwt_token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   3600,
			Secure:   true,
		}
		http.SetCookie(w, &cookie)

	} else {
		log.Printf("Error Login User afeter calling driver: %s /n", err)
		http.Error(w, "User or password are wrong", http.StatusForbidden)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"id": id})

}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("jwt_token")
	cookie.MaxAge = -1
	cookie.Expires = time.Unix(0, 0)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	http.SetCookie(w, cookie)

}

func CheckCookie(w http.ResponseWriter, r *http.Request) {
	id, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}
	username, err := driver.GetUsername(id)

	if err != nil {
		http.Error(w, "user id doesnt match", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"id": strconv.Itoa(id), "username": username})
}
