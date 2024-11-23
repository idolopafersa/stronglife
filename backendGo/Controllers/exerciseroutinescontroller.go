package controllers

import (
	driver "backendgo/Driver"
	"backendgo/security"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetrExercises(w http.ResponseWriter, r *http.Request) {
	routineID := r.URL.Query().Get("routine")

	if routineID == "" {
		http.Error(w, "Routine  is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	exercises, err := driver.GetExercisesInRoutine(routineID)
	if err != nil {
		log.Printf("Error GET exercise in Routine driver: %s /n", err)
		http.Error(w, "Error getting exercises of routine", http.StatusInternalServerError)
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

func DelrExercise(w http.ResponseWriter, r *http.Request) {
	routineID := r.URL.Query().Get("routine")
	exerciseID := r.URL.Query().Get("exercise")

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	erre := driver.RemoveExerciseFromRoutine(routineID, exerciseID)
	if erre != nil {
		log.Printf("Error DEl exercise in Routine driver: %s /n", err)
		http.Error(w, "Error removing exercise from routine", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Exercise removed from routine successfully"})
}

func AddrExercise(w http.ResponseWriter, r *http.Request) {

	routineID := r.URL.Query().Get("routine")
	exerciseID := r.URL.Query().Get("exercise")

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	erre := driver.AddExerciseToRoutine(routineID, exerciseID)
	if erre != nil {
		log.Printf("Error POST exercise in Routine driver: %s /n", err)
		http.Error(w, "Error adding exercise to routine", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Exercise added to routine successfully"})
}
