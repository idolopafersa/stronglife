package controllers

import (
	driver "backendgo/Driver"
	"backendgo/security"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetrExercises(w http.ResponseWriter, r *http.Request) {
	routineID := r.URL.Query().Get("routine_id")

	if routineID == "" {
		http.Error(w, "Routine ID parameter is missing", http.StatusBadRequest)
		return
	}

	exercises, err := driver.GetExercisesInRoutine(routineID)
	if err != nil {
		log.Printf("Error GET exercises in routine driver: %s\n", err)
		http.Error(w, "Error getting exercises in routine", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if len(exercises) > 0 {
		fmt.Print("Hay mas de uno")
		json.NewEncoder(w).Encode(exercises)
	} else {
		fmt.Print("No hay ejercicios")
		w.Write([]byte("[]"))
	}
}

func DelrExercise(w http.ResponseWriter, r *http.Request) {
	routineID := r.URL.Query().Get("routine_id")
	exerciseID := r.URL.Query().Get("exercise_id")

	if routineID == "" || exerciseID == "" {
		http.Error(w, "Routine ID or Exercise ID parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	erre := driver.RemoveExerciseFromRoutine(routineID, exerciseID, userID)
	if erre != nil {
		if erre.Error() == "forbidden: user does not have access to this routine" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error DEL exercise in Routine driver: %s\n", err)
		http.Error(w, "Error removing exercise from routine", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Exercise removed from routine successfully"})
}

func AddrExercise(w http.ResponseWriter, r *http.Request) {
	routineID := r.URL.Query().Get("routine_id")
	exerciseID := r.URL.Query().Get("exercise_id")

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	if routineID == "" || exerciseID == "" {
		http.Error(w, "Routine ID or Exercise ID parameter is missing", http.StatusBadRequest)
		return
	}

	routineIDInt, _ := strconv.Atoi(routineID)
	exerciseIDInt, _ := strconv.Atoi(exerciseID)

	erre := driver.AddExerciseToRoutine(routineIDInt, exerciseIDInt, userID)
	if erre != nil {
		if erre.Error() == "forbidden: user does not have access to this routine" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error POST exercise in Routine driver: %s\n", err)
		http.Error(w, "Error adding exercise to routine", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Exercise added to routine successfully"})
}
