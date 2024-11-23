package controllers

import (
	driver "backendgo/Driver"
	structmodels "backendgo/StructModels"
	"backendgo/security"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetExercise(w http.ResponseWriter, r *http.Request) {
	var exe structmodels.Exercise
	
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	exe, erre := driver.GetExercise(id)
	if erre != nil {
		log.Printf("Error GET exercise driver : %s /n", err)
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
		return
	}

	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	id, err := driver.PostExercise(exercise)
	if err != nil {
		log.Printf("Error POST exercise driver : %s /n", err)
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

	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	erre := driver.DelExercise(id)
	if erre != nil {
		log.Printf("Error DEL exercise driver: %s /n", err)
		http.Error(w, "Exercise not found", http.StatusNotFound)
		return
	}
}
func PutExercise(w http.ResponseWriter, r *http.Request) {
	var exercise structmodels.Exercise

	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
		http.Error(w, "Payload inválido", http.StatusBadRequest)

		return
	}

	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	erre := driver.PutExercise(exercise)
	if erre != nil {
		log.Printf("Error PUT exercise  driver: %s /n", err)
		http.Error(w, "Error al actualizar el ejercicio", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Ejercicio actualizada exitosamente"})
}

func GetAlExercises(w http.ResponseWriter, r *http.Request) {

	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	exercises, erre := driver.GetAlExercises()
	if erre != nil {
		log.Printf("Error GET all exercises  driver: %s /n", err)
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
