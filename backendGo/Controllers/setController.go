package controllers

import (
	driver "backendgo/Driver"
	structmodels "backendgo/StructModels"
	"backendgo/security"
	"encoding/json"
	"log"
	"net/http"
)

func GetSet(w http.ResponseWriter, r *http.Request) {
	routineID := r.URL.Query().Get("routine_id")
	exerciseID := r.URL.Query().Get("exercise_id")
	setNumber := r.URL.Query().Get("set_number")

	if routineID == "" || exerciseID == "" || setNumber == "" {
		http.Error(w, "RoutineId, ExerciseId or SetNumber parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	set, err := driver.GetSet(routineID, exerciseID, setNumber, userID)
	if err != nil {
		if err.Error() == "forbidden: user does not have access to this set" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error GET Set driver: %s\n", err)
		http.Error(w, "Set not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(set)
}

func GetAlSet(w http.ResponseWriter, r *http.Request) {
	routineID := r.URL.Query().Get("routine_id")
	exerciseID := r.URL.Query().Get("exercise_id")

	if routineID == "" || exerciseID == "" {
		http.Error(w, "RoutineId or ExerciseId parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	sets, err := driver.GetAlSet(routineID, exerciseID, userID)
	if err != nil {
		log.Printf("Error GET all sets driver: %s\n", err)
		http.Error(w, "Sets not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sets)
}

func GetAlSetRoutine(w http.ResponseWriter, r *http.Request) {
	routineID := r.URL.Query().Get("routine_id")

	if routineID == "" {
		http.Error(w, "RoutineId parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	sets, err := driver.GetAlSetRoutine(routineID, userID)
	if err != nil {
		log.Printf("Error GET all sets for routine driver: %s\n", err)
		http.Error(w, "Sets not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sets)
}

func PostSet(w http.ResponseWriter, r *http.Request) {
	var set structmodels.Set

	if err := json.NewDecoder(r.Body).Decode(&set); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	// Verificar si el usuario es el dueño de la rutina y el ejercicio
	_, err = driver.GetRoutine(set.RoutineID, userID)
	if err != nil {
		if err.Error() == "forbidden: user does not have access to this routine" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		http.Error(w, "Routine not found", http.StatusNotFound)
		return
	}

	_, err = driver.GetExercise(set.ExerciseID, userID)
	if err != nil {
		if err.Error() == "forbidden: user does not have access to this exercise" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		http.Error(w, "Exercise not found", http.StatusNotFound)
		return
	}

	err = driver.PostSet(set, userID)
	if err != nil {
		log.Printf("Error POST Set driver: %s\n", err)
		http.Error(w, "Error creating set", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Set created"})
}

func PutSet(w http.ResponseWriter, r *http.Request) {
	var set structmodels.Set

	if err := json.NewDecoder(r.Body).Decode(&set); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	err = driver.PutSet(set, userID)
	if err != nil {
		if err.Error() == "forbidden: user does not have access to this set" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error PUT Set driver: %s\n", err)
		http.Error(w, "Error updating set", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Set updated"})
}

func DelSet(w http.ResponseWriter, r *http.Request) {
	routineID := r.URL.Query().Get("routine_id")
	exerciseID := r.URL.Query().Get("exercise_id")
	setNumber := r.URL.Query().Get("set_number")

	if routineID == "" || exerciseID == "" || setNumber == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	err = driver.DelSet(routineID, exerciseID, setNumber, userID)
	if err != nil {
		if err.Error() == "forbidden: user does not have access to this set" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error DEL Set driver: %s\n", err)
		http.Error(w, "Error deleting set", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
