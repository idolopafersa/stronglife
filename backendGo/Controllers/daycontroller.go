package controllers

import (
	driver "backendgo/Driver"
	"backendgo/security"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Getday(w http.ResponseWriter, r *http.Request) {

	date := r.URL.Query().Get("date")

	if date == "" {
		http.Error(w, "date parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusForbidden)
		return
	}

	day, err := driver.GetDay(userID, date)
	if err != nil {
		log.Printf("Error getting day driver: %s /n", err)
		http.Error(w, "Could not get day", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(day)
}

func AddMealToDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	mealID := r.URL.Query().Get("meal_id")

	if date == "" || mealID == "" {
		http.Error(w, "date  or meal_id parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
		return
	}

	erre := driver.AddMealToDay(date, strconv.Itoa(userID), mealID)
	if erre != nil {
		log.Printf("Error Adding Meal to Day day driver: %s /n", err)
		http.Error(w, "Could not add meal to day", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Meal added successfully")

}

func AddRoutineToDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	routineID := r.URL.Query().Get("routine_id")

	if date == "" || routineID == "" {
		http.Error(w, "date  or routine_id parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
		return
	}

	erre := driver.AddRoutineToDay(date, strconv.Itoa(userID), routineID)
	if erre != nil {
		log.Printf("Error adding Routine to Day  driver: %s /n", err)
		http.Error(w, "Could not add routine to day", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("routine added successfully")

}

func DelMealToDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	mealID := r.URL.Query().Get("meal_id")

	if date == "" || mealID == "" {
		http.Error(w, "date  or meal_id parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
		return
	}

	erre := driver.DelMealToDay(date, strconv.Itoa(userID), mealID)
	if erre != nil {
		log.Printf("Error DEL meal to  day driver: %s /n", err)
		http.Error(w, "Could not add meal to day", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Meal deleted successfully")

}
