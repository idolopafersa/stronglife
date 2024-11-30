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

func PostRoutine(w http.ResponseWriter, r *http.Request) {
	var routine structmodels.NewRoutine

	if err := json.NewDecoder(r.Body).Decode(&routine); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	id, err := driver.PostRoutine(routine, userID)
	if err != nil {
		log.Printf("Error POST Routine driver: %s\n", err)
		http.Error(w, "Error creating routine", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func GetRoutine(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("routine_id")

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
	routine, err := driver.GetRoutine(intID, userID)
	if err != nil {
		if err.Error() == "forbidden: user does not have access to this routine" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error GET Routine driver: %s\n", err)
		http.Error(w, "Routine not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routine)
}

func PutRoutine(w http.ResponseWriter, r *http.Request) {
	var routine structmodels.Routine

	if err := json.NewDecoder(r.Body).Decode(&routine); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	err = driver.PutRoutine(routine, userID)
	if err != nil {
		if err.Error() == "forbidden: user does not have access to this routine" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error PUT Routine driver: %s\n", err)
		http.Error(w, "Error updating routine", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Routine updated"})
}

func DelRoutine(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("routine_id")

	if id == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	erre := driver.DelRoutine(id, userID)
	if erre != nil {
		if erre.Error() == "forbidden: user does not have access to this routine" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Printf("Error DEL Routine driver: %s\n", err)
		http.Error(w, "Routine not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func AllRoutines(w http.ResponseWriter, r *http.Request) {

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	routines, erre := driver.GetAlRoutines(userID)
	if erre != nil {
		log.Printf("Error GETALL  Routine driver: %s /n", err)
		http.Error(w, "routine not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routines)

}
