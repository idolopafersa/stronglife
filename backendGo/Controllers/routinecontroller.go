package controllers

import (
	driver "backendgo/Driver"
	structmodels "backendgo/StructModels"
	"backendgo/security"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostRoutine(w http.ResponseWriter, r *http.Request) {
	var routine structmodels.NewRoutine

	if err := json.NewDecoder(r.Body).Decode(&routine); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		fmt.Println(err)
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

	id, err := driver.PostRoutine(routine)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error creating routine", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func GetRoutine(w http.ResponseWriter, r *http.Request) {
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
	routine, err := driver.GetRoutine(id)
	if err != nil {
		fmt.Println(err)
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

	errw := driver.PutRoutine(routine)
	if errw != nil {
		fmt.Println(errw)
		http.Error(w, "Error updating routine", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Routine updated "})
}

func DelRoutine(w http.ResponseWriter, r *http.Request) {

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

	erre := driver.DelRoutine(id)
	if erre != nil {
		fmt.Print(erre)
		http.Error(w, "routine not found", http.StatusNotFound)
		return
	}
}

func AllRoutines(w http.ResponseWriter, r *http.Request) {

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

	routines, erre := driver.GetAlRoutines()
	if erre != nil {
		fmt.Print(erre)
		http.Error(w, "routine not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routines)

}
