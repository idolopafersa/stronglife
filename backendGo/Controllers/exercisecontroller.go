package controllers

import (
	driver "backendgo/Driver"
	structmodels "backendgo/StructModels"
	"backendgo/security"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetExercise(w http.ResponseWriter, r *http.Request) {
	var exe structmodels.Exercise

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Name parameter is missing", http.StatusBadRequest)
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

	exe, erre := driver.GetExercise(id)
	if erre != nil {
		fmt.Print(erre)
		http.Error(w, "Exercise not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exe)
}

func PostExercise(w http.ResponseWriter, r *http.Request) {
	var exercise structmodels.NewExercise

	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
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

	id, err := driver.PostExercise(exercise)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Error adding exercise", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func DelExercise(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {

		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
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

	erre := driver.DelExercise(id)
	if erre != nil {
		fmt.Print(err)
		http.Error(w, "Exercise not found", http.StatusNotFound)
		return
	}
}
func PutExercise(w http.ResponseWriter, r *http.Request) {
	var exercise structmodels.Exercise

	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
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

	erre := driver.PutExercise(exercise)
	fmt.Print(erre)
	if erre != nil {
		fmt.Print(erre)
		http.Error(w, "Error al actualizar el ejercicio", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Ejercicio actualizada exitosamente"})
}

func GetAlExercises(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
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

	exercises, erre := driver.GetAlExercises()
	if erre != nil {
		fmt.Print(erre)
		http.Error(w, "routine not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if len(exercises) > 0 {
		fmt.Print("Hay mas de uno")
		json.NewEncoder(w).Encode(exercises)
	} else {
		fmt.Print("Hay mas de uno")
		w.Write([]byte("[]"))
	}

}
