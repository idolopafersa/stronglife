package driver

import (
	structmodels "backendgo/StructModels"
	"fmt"
)

func GetSet(routineID, exerciseID string) (structmodels.Set, error) {
	// Uso LIMIT 1 para devolver unicamente un valor, despues habria que hacer un order
	// habra que asignar un usuario o verificar de que rutina se pide (ya que va asociado a un user)

	var Set structmodels.Set //estructura set
	//query_user := "SELECT id FROM Routines WHERE user_id = ?"
	//err := db.QueryRow(query_user, routineID).Scan(&Set.RoutineID)

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

func PutSet(RoutineID, ExerciseID, Set_number, Weight, Reps string, userID int) error {
	query := `UPDATE Sets SET reps = ?, weight = ? WHERE routine_id = ? AND exercise_id = ? AND set_number = ?`
	_, err := db.Exec(query, Reps, Weight, RoutineID, ExerciseID, Set_number)

	if err != nil {
		fmt.Println("Error updating set:", err)
		return err
	}
	return nil
}

func DelSet(RoutineID, ExerciseID, Set_number string, userID int) error {
	var setID int
	queryGetID := "SELECT id FROM Sets WHERE routine_id = ? AND exercise_id = ? AND set_number = ?"
	db.QueryRow(queryGetID, RoutineID, ExerciseID, Set_number).Scan(&setID)

	query := "DELETE FROM Sets WHERE id = ?;"
	_, err := db.Exec(query, setID)
	fmt.Print(err)

	return err
}
