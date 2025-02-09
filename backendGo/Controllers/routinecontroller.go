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
	var routine structmodels.Routine

	// Decode the incoming request body into the Routine struct
	if err := json.NewDecoder(r.Body).Decode(&routine); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Verify the user ID from the cookie
	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	// Create the routine in the database
	id, err := driver.PostRoutine(routine, userID)
	if err != nil {
		log.Printf("Error POST Routine driver: %s\n", err)
		http.Error(w, "Error creating routine", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
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

// way to create a routine with all the exercises and sets, if the given routine already exists, itll edit the elements
func Training(w http.ResponseWriter, r *http.Request) {
	var routine structmodels.Routine

	// Decode the incoming request body into the Routine struct
	if err := json.NewDecoder(r.Body).Decode(&routine); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Verify the user ID from the cookie
	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error verifying cookie", http.StatusUnauthorized)
		return
	}

	// Check if the exercises belong to the user
	for _, exercise := range routine.Exercises {
		_, err := driver.GetExercise(exercise.ID, userID)
		if err != nil {
			log.Printf("Failed to get exercise: %d", exercise.ID)
			http.Error(w, "You aren't the owner of the exercise or it doesn't exist", http.StatusForbidden)
			return
		}
	}

	// Create the routine in the database
	id, err := driver.PostRoutine(routine, userID)
	if err != nil {
		log.Printf("Error POST Routine driver: %s\n", err)
		http.Error(w, "Error creating routine", http.StatusInternalServerError)
		return
	}

	var response struct {
		ID        int    `json:"routine_id"`
		Name      string `json:"routine_name"`
		Exercises []struct {
			ID          int                `json:"exercise_id"`
			Name        string             `json:"name"`
			Description string             `json:"description"`
			Sets        []structmodels.Set `json:"sets"`
		} `json:"exercises"`
	}

	// Add exercises to the routine and create sets
	for _, exercise := range routine.Exercises {
		err := driver.AddExerciseToRoutine(id, exercise.ID, userID)
		if err != nil {
			log.Printf("Failed to add exercise: %d", exercise.ID)
			http.Error(w, "You aren't the owner of the exercise or it doesn't exist", http.StatusForbidden)
			return
		}

		var exerciseWithSets struct {
			ID          int                `json:"exercise_id"`
			Name        string             `json:"name"`
			Description string             `json:"description"`
			Sets        []structmodels.Set `json:"sets"`
		}
		exerciseWithSets.ID = exercise.ID
		exerciseWithSets.Name = exercise.Name
		exerciseWithSets.Description = exercise.Description

		for _, set := range exercise.Sets {
			newSet := structmodels.Set{
				RoutineID:  id,
				ExerciseID: exercise.ID,
				SetNumber:  set.SetNumber,
				Reps:       set.Reps,
				Weight:     set.Weight,
			}
			err := driver.PostSet(newSet, userID) // Handle error if needed
			if err != nil {
				log.Printf("Error adding set: %v", err)
				http.Error(w, "Error adding set", http.StatusInternalServerError)
				return
			}

			// Append the newly created set to the exercise's sets
			exerciseWithSets.Sets = append(exerciseWithSets.Sets, newSet)
		}

		// Add the exercise (with sets) to the response
		response.Exercises = append(response.Exercises, exerciseWithSets)
	}

	// Return the created routine as a JSON response, including sets
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Status code 201 - Created

	// Encode and send the JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %s", err)
		http.Error(w, "Error returning JSON response", http.StatusInternalServerError)
	}
}
