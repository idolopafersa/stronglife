package controllers

import (
	driver "backendgo/Driver"
	structmodels "backendgo/StructModels"
	"backendgo/security"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetExercise(w http.ResponseWriter, r *http.Request) {
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
	intID, _ := strconv.Atoi(id)
	exe, erre := driver.GetExercise(intID, userID)
	if erre != nil {
		if erre.Error() == "forbidden: user does not have access to this exercise" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error GET exercise driver: %s\n", erre)
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

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	id, err := driver.PostExercise(exercise, userID)
	if err != nil {
		log.Printf("Error POST exercise driver: %s\n", err)
		http.Error(w, "Error adding exercise", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func DelExercise(w http.ResponseWriter, r *http.Request) {
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
	intID, _ := strconv.Atoi(id)

	erre := driver.DelExercise(intID, userID)

	if erre != nil {
		if erre.Error() == "forbidden: user does not have access to this exercise" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error DEL exercise driver: %s\n", erre)
		http.Error(w, "Exercise not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func PutExercise(w http.ResponseWriter, r *http.Request) {
	var exercise structmodels.Exercise

	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	erre := driver.PutExercise(exercise, userID)
	if erre != nil {
		if erre.Error() == "forbidden: user does not have access to this exercise" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error PUT exercise driver: %s\n", erre)
		http.Error(w, "Error updating exercise", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Exercise updated successfully"})
}

func GetAlExercises(w http.ResponseWriter, r *http.Request) {

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	exercises, erre := driver.GetAlExercises(userID)
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
