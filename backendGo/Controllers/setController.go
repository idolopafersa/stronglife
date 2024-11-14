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

func GetSet(w http.ResponseWriter, r *http.Request) { //Devuelve un único set, dado el id de la rutina y el ejercicio
	RoutineId := r.URL.Query().Get("RoutineId")
	ExerciseId := r.URL.Query().Get("ExerciseId")

	_, err := security.VerifyCookie(r)
	if err != nil {
		http.Error(w, "Error Cookie", http.StatusNotFound)
	}

	set, err := driver.GetSet()
	if err != nil {
		log.Printf("Error Getting Sets in driver: %s /n", err)
		http.Error(w, "Couldnt get sets", 404)
		return
	}

	json.NewEncoder(w).Encode(sets)
}
