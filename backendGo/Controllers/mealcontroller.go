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
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	meal, err := driver.GetMeal(id, userID)

	if err != nil {
		log.Printf("Error GET Meals driver: %s /n", err)
		http.Error(w, "Meal not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(meal)

}

func PostMeal(w http.ResponseWriter, r *http.Request) {
	var nmeal structmodels.NewMeal

	if err := json.NewDecoder(r.Body).Decode(&nmeal); err != nil { //Recogemos el body, debe haber todo un Meal
		http.Error(w, "Invalid payload", 400)
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	id, err := driver.PostMeal(nmeal, userID)
	if err != nil {
		log.Printf("Error POST Meals driver: %s /n", err)
		http.Error(w, "couldnt Post", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func DelMeal(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	erre := driver.DelMeal(id)
	if erre != nil {
		log.Printf("Error DEL Meals driver: %s /n", erre)
		http.Error(w, "No Meal", http.StatusNotFound)
	}
}

func PutMeal(w http.ResponseWriter, r *http.Request) {
	var meal structmodels.Meal

	if err := json.NewDecoder(r.Body).Decode(&meal); err != nil {
		http.Error(w, "Payload inválido", http.StatusBadRequest)
		fmt.Print(err)
		return
	}

	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	erre := driver.UpdateMeal(meal)
	if erre != nil {
		log.Printf("Error PUT Meals driver: %s /n", erre)
		http.Error(w, "Error al actualizar la comida", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Comida actualizada exitosamente"})
}

func AllMeals(w http.ResponseWriter, r *http.Request) {

	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	meals, erre := driver.GetAllMeals()
	if erre != nil {
		log.Printf("Error GET AllMeals driver: %s /n", erre)
		http.Error(w, "meal not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meals)

}
