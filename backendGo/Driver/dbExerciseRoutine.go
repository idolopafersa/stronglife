package driver

import structmodels "backendgo/StructModels"

func AddExerciseToRoutine(routineID, exerciseID string) error {
	err := db.QueryRow("SELECT name FROM Routines WHERE id = ? AND user_id = ?", routineID, userID)

	if err == sql.ErrNoRows {
		fmt.Println("No routine found with that ID for the current user")
	} else if err != nil {
		log.Fatalf("Query error: %v", err)
	} else {
		fmt.Printf("Routine found: %s\n", name)
		query := "INSERT INTO RoutineExercises (routine_id, exercise_id) VALUES (?, ?)"
		_, err := db.Exec(query, routineID, exerciseID)
	}

	return err
}

func RemoveExerciseFromRoutine(routineID, exerciseID string) error {
	// Primero comprobamos que la rutina existe y pertenece al usuario
	err := db.QueryRow("SELECT name FROM Routines WHERE id = ? AND user_id = ?", routineID, userID)

	//Comprobamos que existe la vinculación entre la rutina y ejercicio (RoutineExercises)
	err:= db.QueryRow("SELECT * FROM RoutineExercises WHERE routine_id = ? AND exercise_id = ?", routineID, exerciseID)
	
	if err == sql.ErrNoRows {
		fmt.Println("No exercise found with that ID on that rutine for the current user")
	} else if err != nil {
		log.Fatalf("Query error: %v", err)
	} else {
		fmt.Printf("Exercise found: %s\n", name)
		query := "DELETE FROM RoutineExercises WHERE routine_id = ? AND exercise_id = ?"
		_, err := db.Exec(query, routineID, exerciseID)
	}

	return err
}

func GetExercisesInRoutine(routineID string) ([]structmodels.Exercise, error) {
	//Falta por verificar seguridad
	query := `SELECT e.* FROM Exercises e
              INNER JOIN RoutineExercises re ON e.id = re.exercise_id
              WHERE re.routine_id?`

	rows, err := db.Query(query, routineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []structmodels.Exercise
	for rows.Next() {
		var exercise structmodels.Exercise
		if err := rows.Scan(&exercise.ID, &exercise.Name, &exercise.Description); err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}

	return exercises, nil
}
