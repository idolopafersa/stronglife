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
	query := "SELECT id, routine_id, exercise_id, set_number, reps, weight FROM Sets WHERE routine_id = ? AND exercise_id = ? LIMIT 1"
	err := db.QueryRow(query, routineID, exerciseID).Scan(&Set.Set_number, &Set.Reps, &Set.Weight, &Set.RoutineID, &Set.ExerciseID)

	if err != nil {
		fmt.Println("Error getting set", err)
		return Set, err
	}
	return Set, nil
}

func GetAlSet(routineID, exerciseID string) ([]structmodels.Set, error) {

	var Sets []structmodels.Set
	query := fmt.Sprintf("SELECT id, routine_id, exercise_id, set_number, reps, weight FROM Sets WHERE routine_id = %s AND exercise_id = %s;", routineID, exerciseID)
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Set structmodels.Set
		if err := rows.Scan(&Set.ID, &Set.RoutineID, &Set.ExerciseID, &Set.Set_number, &Set.Reps, &Set.Weight); err != nil {
			return nil, err
		}
		Sets = append(Sets, Set)
	}

	return Sets, nil
}

func GetAlSetRoutine(routineID string) ([]structmodels.Set, error) {

	var Sets []structmodels.Set
	query := fmt.Sprintf("SELECT id, routine_id, exercise_id, set_number, reps, weight FROM Sets WHERE routine_id = %s;", routineID)
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Set structmodels.Set
		if err := rows.Scan(&Set.ID, &Set.RoutineID, &Set.ExerciseID, &Set.Set_number, &Set.Reps, &Set.Weight); err != nil {
			return nil, err
		}
		Sets = append(Sets, Set)
	}

	return Sets, nil
}
