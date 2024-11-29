package driver

import (
	structmodels "backendgo/StructModels"
	"fmt"
)

func GetSet(routineID, exerciseID string, userID int) (structmodels.Set, error) {
	// Uso LIMIT 1 para devolver unicamente un valor, despues habria que hacer un order
	// habra que asignar un usuario o verificar de que rutina se pide (ya que va asociado a un user)

	var Set structmodels.Set //estructura set
	//query_user := "SELECT id FROM Routines WHERE user_id = ?"
	//err := db.QueryRow(query_user, routineID).Scan(&Set.RoutineID)

	query := fmt.Sprintf("SELECT  routine_id, exercise_id, set_number, reps, weight FROM Sets WHERE routine_id = ? AND exercise_id = ? AND user_id = %d LIMIT 1", userID)
	err := db.QueryRow(query, routineID, exerciseID).Scan(&Set.Set_number, &Set.Reps, &Set.Weight, &Set.RoutineID, &Set.ExerciseID)

	if err != nil {
		fmt.Println("Error getting set", err)
		return Set, err
	}
	return Set, nil
}

func GetAlSet(routineID, exerciseID string, userID int) ([]structmodels.Set, error) {

	var Sets []structmodels.Set
	query := fmt.Sprintf("SELECT  routine_id, exercise_id, set_number, reps, weight FROM Sets WHERE routine_id = %s AND exercise_id = %s AND user_id = %d ;", routineID, exerciseID, userID)
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Set structmodels.Set
		if err := rows.Scan(&Set.RoutineID, &Set.ExerciseID, &Set.Set_number, &Set.Reps, &Set.Weight); err != nil {
			return nil, err
		}
		Sets = append(Sets, Set)
	}

	return Sets, nil
}

func GetAlSetRoutine(routineID string, userID int) ([]structmodels.Set, error) {

	var Sets []structmodels.Set
	query := fmt.Sprintf("SELECT  routine_id, exercise_id, set_number, reps, weight FROM Sets WHERE routine_id = %s; AND user_id = %d", routineID, userID)
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Set structmodels.Set
		if err := rows.Scan(&Set.RoutineID, &Set.ExerciseID, &Set.Set_number, &Set.Reps, &Set.Weight); err != nil {
			return nil, err
		}
		Sets = append(Sets, Set)
	}

	return Sets, nil
}

func PutSet(Set structmodels.Set, userID int) error {
	query := fmt.Sprintf(`UPDATE Sets SET reps = ?, weight = ? WHERE routine_id = ? AND exercise_id = ? AND user_id= %d AND set_number = ? AND id = ?`, userID)
	_, err := db.Exec(query, Set.Reps, Set.Weight, Set.RoutineID, Set.ExerciseID, Set.Set_number)

	if err != nil {
		fmt.Println("Error updating set:", err)
		return err
	}
	return nil
}

func PostSet(Set structmodels.Set, userID int) error {
	query := `
        INSERT INTO Sets (routine_id, exercise_id, set_number, weight, reps, user_id) 
        VALUES (?, ?, ?, ?, ?, ?)
    `
	_, err := db.Exec(query, Set.RoutineID, Set.ExerciseID, Set.Set_number, Set.Weight, Set.Reps, userID)
	if err != nil {
		fmt.Println("Error inserting set:", err)
		return err
	}
	return nil
}

func DelSet(RoutineID, ExerciseID, Set_number string, userID int) error {
	var setID int
	var dbUserID int
	queryGetID := "SELECT id , user_id FROM Sets WHERE routine_id = ? AND exercise_id = ? AND set_number = ?"
	db.QueryRow(queryGetID, RoutineID, ExerciseID, Set_number).Scan(&setID, &dbUserID)

	if dbUserID != userID {
		return fmt.Errorf("userID does not match the owner of the set")
	}

	query := "DELETE FROM Sets WHERE id = ?;"
	_, err := db.Exec(query, setID)
	fmt.Print(err)

	return err
}
