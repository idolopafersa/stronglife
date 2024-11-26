package driver

import (
	structmodels "backendgo/StructModels"
	"fmt"
)

func GetSet(routineID, exerciseID string) (structmodels.Set, error) {
	// Uso top para devolver unicamente un valor, despues habria que hacer un order
	// habra que asignar un usuario o verificar de que rutina se pide (ya que va asociado a un user)

	var Set structmodels.Set //estructura set

	// WORKING
	// Falta la seguridad
	// hace falta id para el set? -> Ahora mismo no existe
	// Pero si existe en struct, corregir struct o bdd
	query := "SELECT set_number, reps, weight, routine_id, exercise_id FROM Sets WHERE routine_id = ? AND exercise_id = ? LIMIT 1"
	err := db.QueryRow(query, routineID, exerciseID).Scan(&Set.Set_number, &Set.Reps, &Set.Weight, &Set.RoutineID, &Set.ExerciseID)

	if err != nil {
		fmt.Println("Error getting routine:", err)
		return Set, err
	}
	return Set, nil
}
