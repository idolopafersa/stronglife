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

	//Ahora mismo los sets no van a tener seguridad
	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

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
