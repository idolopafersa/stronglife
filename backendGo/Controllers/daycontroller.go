package controllers

import (
	driver "backendgo/Driver"
	"backendgo/security"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Getday(w http.ResponseWriter, r *http.Request) {

	date := r.URL.Query().Get("date")

	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized: No valid cookie", http.StatusUnauthorized)
		return
	}

	// Extract the JWT from the cookie value
	jwtToken := cookie.Value

	if err := security.VerifyToken(jwtToken); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JWT", http.StatusForbidden)
		return
	}
	name, _ := security.ExtractUser(jwtToken)
	fmt.Print("username is" + name + "   ")
	userID, _ := driver.GetUserID(name)

	day, err := driver.GetDay(userID, date)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Could not get day", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(day)
}

func AddMealToDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	mealID := r.URL.Query().Get("meal")

	cookie, err := r.Cookie("token")
	if err != nil {

		http.Error(w, "Unauthorized: No valid cookie", http.StatusUnauthorized)
		return
	}

	// Extract the JWT from the cookie value
	jwtToken := cookie.Value

	if err := security.VerifyToken(jwtToken); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JWT", http.StatusForbidden)
		return
	}

	name, _ := security.ExtractUser(jwtToken)
	userID, _ := driver.GetUserID(name)

	erre := driver.AddMealToDay(date, strconv.Itoa(userID), mealID)
	if erre != nil {
		fmt.Print(erre)
		http.Error(w, "Could not add meal to day", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Meal added successfully")

}

func AddRoutineToDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	routineID := r.URL.Query().Get("routine")

	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized: No valid cookie", http.StatusUnauthorized)
		return
	}

	// Extract the JWT from the cookie value
	jwtToken := cookie.Value

	if err := security.VerifyToken(jwtToken); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JWT", http.StatusForbidden)
		return
	}
	name, _ := security.ExtractUser(jwtToken)
	userID, _ := driver.GetUserID(name)

	erre := driver.AddRoutineToDay(date, strconv.Itoa(userID), routineID)
	if erre != nil {
		fmt.Print(erre)
		http.Error(w, "Could not add routine to day", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("routine added successfully")

}
func DelMealToDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	mealID := r.URL.Query().Get("meal")

	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized: No valid cookie", http.StatusUnauthorized)
		return
	}

	// Extract the JWT from the cookie value
	jwtToken := cookie.Value

	if err := security.VerifyToken(jwtToken); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JWT", http.StatusForbidden)
		return
	}

	name, _ := security.ExtractUser(jwtToken)
	userID, _ := driver.GetUserID(name)

	erre := driver.DelMealToDay(date, strconv.Itoa(userID), mealID)
	if erre != nil {
		http.Error(w, "Could not add meal to day", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Meal deleted successfully")

}
