package controllers

import (
	driver "backendgo/Driver"
	"backendgo/security"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetrExercises(w http.ResponseWriter, r *http.Request) {
	routineID := r.URL.Query().Get("routine")

	if routineID == "" {
		http.Error(w, "Routine  is missing", http.StatusBadRequest)
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

	exercises, err := driver.GetExercisesInRoutine(routineID)
	if err != nil {
		fmt.Println(err)
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

	erre := driver.RemoveExerciseFromRoutine(routineID, exerciseID)
	if erre != nil {
		fmt.Println(err)
		http.Error(w, "Error removing exercise from routine", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Exercise removed from routine successfully"})
}

func AddrExercise(w http.ResponseWriter, r *http.Request) {

	routineID := r.URL.Query().Get("routine")
	exerciseID := r.URL.Query().Get("exercise")

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

	erre := driver.AddExerciseToRoutine(routineID, exerciseID)
	if erre != nil {
		fmt.Println(err)
		http.Error(w, "Error adding exercise to routine", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Exercise added to routine successfully"})
}
