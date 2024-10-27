package driver

import structmodels "backendgo/StructModels"

func AddExerciseToRoutine(routineID, exerciseID string) error {
	query := "INSERT INTO RoutineExercises (routine_id, exercise_id) VALUES (?, ?)"
	_, err := db.Exec(query, routineID, exerciseID)
	return err
}

func RemoveExerciseFromRoutine(routineID, exerciseID string) error {
	query := "DELETE FROM RoutineExercises WHERE routine_id = ? AND exercise_id = ?"
	_, err := db.Exec(query, routineID, exerciseID)
	return err
}

func GetExercisesInRoutine(routineID string) ([]structmodels.Exercise, error) {
	query := `SELECT e.* FROM Exercises e
              INNER JOIN RoutineExercises re ON e.id = re.exercise_id
              WHERE re.routine_id = ?`

	rows, err := db.Query(query, routineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []structmodels.Exercise
	for rows.Next() {
		var exercise structmodels.Exercise
		if err := rows.Scan(&exercise.ID, &exercise.Name, &exercise.Sets, &exercise.Repetitions, &exercise.Description, &exercise.PhotoURL); err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}

	return exercises, nil
}
