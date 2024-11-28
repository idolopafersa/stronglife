package controllers

import (
	driver "backendgo/Driver"
	"backendgo/security"
	"encoding/json"
	"log"
	"net/http"
)

//hay que hacer las funciones pequeñas por si se necesitan de forma puntual. Pero hay que hacer tanto un Metodo Get y Post,
// que cogan toda la información de un ejercicio o Rutina, para guardarla en caché en el móvil y sea más rápido.
// Seguramente solo lo vea Ale, así que, que tal guapo.
// Para guapo tú ;)

func GetSet(w http.ResponseWriter, r *http.Request) { //Devuelve un único set, dado el id de la rutina y el ejercicio
	RoutineId := r.URL.Query().Get("RoutineId")
	ExerciseId := r.URL.Query().Get("ExerciseId")

	// id NULL
	if RoutineId == "" || ExerciseId == "" {
		http.Error(w, "RoutineId or ExerciseId parameter is missing", http.StatusBadRequest)
		return
	}

	//Ahora mismo los sets no van a tener seguridad con el userID
	//userID, err := security.VerifyCookie(r)
	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	//	set, err := driver.GetSet(RoutineId, ExerciseId, userID)
	set, err := driver.GetSet(RoutineId, ExerciseId)
	if err != nil {
		log.Printf("Error Getting Sets in driver: %s /n", err)
		http.Error(w, "Couldnt get sets", 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(set)
}

func GetAlSet(w http.ResponseWriter, r *http.Request) {
	RoutineId := r.URL.Query().Get("RoutineId")
	ExerciseId := r.URL.Query().Get("ExerciseId")

	// id NULL
	if RoutineId == "" || ExerciseId == "" {
		http.Error(w, "RoutineId or ExerciseId parameter is missing", http.StatusBadRequest)
		return
	}

	//Ahora mismo los sets no van a tener seguridad
	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	set, err := driver.GetAlSet(RoutineId, ExerciseId)
	if err != nil {
		log.Printf("Error Getting Sets in driver: %s /n", err)
		http.Error(w, "Couldnt get sets", 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(set)
}

func GetAlSetRoutine(w http.ResponseWriter, r *http.Request) {
	RoutineId := r.URL.Query().Get("RoutineId")

	// id NULL
	if RoutineId == "" {
		http.Error(w, "RoutineId parameter is missing", http.StatusBadRequest)
		return
	}

	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	set, err := driver.GetAlSetRoutine(RoutineId)
	if err != nil {
		log.Printf("Error Getting Sets in driver: %s /n", err)
		http.Error(w, "Couldnt get sets", 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(set)
}

func PutSet(w http.ResponseWriter, r *http.Request) {
	RoutineID := r.URL.Query().Get("RoutineID")
	ExerciseID := r.URL.Query().Get("ExerciseID")
	Set_number := r.URL.Query().Get("Set_number")
	Weight := r.URL.Query().Get("Weight")
	Reps := r.URL.Query().Get("Reps")

	if RoutineID == "" {
		http.Error(w, "RoutineID parameter is missing", http.StatusBadRequest)
		return
	} else if ExerciseID == "" {
		http.Error(w, "ExerciseID parameter is missing", http.StatusBadRequest)
		return
	} else if Set_number == "" {
		http.Error(w, "Set_number parameter is missing", http.StatusBadRequest)
		return
	} else if Weight == "" {
		http.Error(w, "Weight parameter is missing", http.StatusBadRequest)
		return
	} else if Reps == "" {
		http.Error(w, "Reps parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	err = driver.PutSet(RoutineID, ExerciseID, Set_number, Weight, Reps, userID)
	if err != nil {
		log.Printf("Error PUT Set driver: %s /n", err)
		http.Error(w, "Error updating set", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Set updated "})
}

func DelSet(w http.ResponseWriter, r *http.Request) {
	RoutineID := r.URL.Query().Get("RoutineID")
	ExerciseID := r.URL.Query().Get("ExerciseID")
	Set_number := r.URL.Query().Get("Set_number")

	if RoutineID == "" {
		http.Error(w, "RoutineID parameter is missing", http.StatusBadRequest)
		return
	} else if ExerciseID == "" {
		http.Error(w, "ExerciseID parameter is missing", http.StatusBadRequest)
		return
	} else if Set_number == "" {
		http.Error(w, "Set_number parameter is missing", http.StatusBadRequest)
		return
	}

	userID, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	erre := driver.DelSet(RoutineID, ExerciseID, Set_number, userID)
	if erre != nil {
		log.Printf("Error DEL Set driver: %s /n", err)
		http.Error(w, "set not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Set deleted successfully"})
}
