package controllers

import (
	driver "backendgo/Driver"
	structmodels "backendgo/StructModels"
	"backendgo/security"
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

	fmt.Print(id)
	meal, err := driver.GetMeal(id)
	fmt.Print(err)
	if err != nil {

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

	id, err := driver.PostMeal(nmeal)
	if err != nil {
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

	fmt.Print(id)

	erre := driver.DelMeal(id)
	if erre != nil {
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

	erre := driver.UpdateMeal(meal)
	fmt.Print(err)
	if erre != nil {
		http.Error(w, "Error al actualizar la comida", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Comida actualizada exitosamente"})
}

func AllMeals(w http.ResponseWriter, r *http.Request) {

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

	meals, erre := driver.GetAllMeals()
	if erre != nil {
		fmt.Print(erre)
		http.Error(w, "meal not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meals)

}
