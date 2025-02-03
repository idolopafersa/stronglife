package controllers

import (
	driver "backendgo/Driver"
	structmodels "backendgo/StructModels"
	"log"

	security "backendgo/security"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetMeal(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	meal, erre := driver.GetMeal(id, userID)
	if erre != nil {
		if erre.Error() == "forbidden: user does not have access to this meal" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error GET meal driver: %s\n", erre)
		http.Error(w, "Meal not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meal)
}

func PostMeal(w http.ResponseWriter, r *http.Request) {
	var meal structmodels.NewMeal

	if err := json.NewDecoder(r.Body).Decode(&meal); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	id, err := driver.PostMeal(meal, userID)
	if err != nil {
		log.Printf("Error POST meal driver: %s\n", err)
		http.Error(w, "Error adding meal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func DelMeal(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	erre := driver.DelMeal(id, userID)
	if erre != nil {
		if erre.Error() == "forbidden: user does not have access to this meal" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error DEL meal driver: %s\n", erre)
		http.Error(w, "Meal not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func PutMeal(w http.ResponseWriter, r *http.Request) {
	var meal structmodels.Meal

	if err := json.NewDecoder(r.Body).Decode(&meal); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	erre := driver.UpdateMeal(meal, userID)
	if erre != nil {
		if erre.Error() == "forbidden: user does not have access to this meal" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error PUT meal driver: %s\n", erre)
		http.Error(w, "Error updating meal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Meal updated successfully"})
}

func AllMeals(w http.ResponseWriter, r *http.Request) {

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	meals, erre := driver.GetAllMeals(userID)
	if erre != nil {
		log.Printf("Error GET AllMeals driver: %s /n", erre)
		http.Error(w, "meal not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meals)

}
