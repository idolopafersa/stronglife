package driver

import (
	structmodels "backendgo/StructModels"
	"database/sql"
	"fmt"
)

func AddExerciseToRoutine(routineID, exerciseID string, userID int) error {
	var name string

	// Comprobar que la rutina pertenece al usuario
	err := db.QueryRow("SELECT name FROM Routines WHERE id = ? AND user_id = ?", routineID, userID).Scan(&name)
	if err == sql.ErrNoRows {
		fmt.Println("No routine found with that ID for the current user")
		return fmt.Errorf("forbidden: user does not have access to this routine")
	} else if err != nil {
		return fmt.Errorf("query error: %v", err)
	}

	// Insertar el ejercicio en la rutina
	query := "INSERT INTO RoutineExercises (routine_id, exercise_id) VALUES (?, ?)"
	_, err = db.Exec(query, routineID, exerciseID)
	if err != nil {
		return fmt.Errorf("failed to add exercise to routine: %v", err)
	}

	fmt.Printf("Exercise added to routine: %s\n", name)
	return nil
}

func RemoveExerciseFromRoutine(routineID, exerciseID string, userID int) error {
	var name string

	// Comprobar que la rutina pertenece al usuario
	err := db.QueryRow("SELECT name FROM Routines WHERE id = ? AND user_id = ?", routineID, userID).Scan(&name)
	if err == sql.ErrNoRows {
		fmt.Println("No routine found with that ID for the current user")
		return fmt.Errorf("forbidden: user does not have access to this routine")
	} else if err != nil {
		return fmt.Errorf("query error: %v", err)
	}

	// Comprobar que el ejercicio está vinculado a la rutina
	err = db.QueryRow("SELECT 1 FROM RoutineExercises WHERE routine_id = ? AND exercise_id = ?", routineID, exerciseID).Scan(new(int))
	if err == sql.ErrNoRows {
		fmt.Println("No exercise found with that ID in the routine for the current user")
		return fmt.Errorf("exercise not found in routine")
	} else if err != nil {
		return fmt.Errorf("query error: %v", err)
	}

	// Eliminar el ejercicio de la rutina
	query := "DELETE FROM RoutineExercises WHERE routine_id = ? AND exercise_id = ?"
	_, err = db.Exec(query, routineID, exerciseID)
	if err != nil {
		return fmt.Errorf("failed to remove exercise from routine: %v", err)
	}

	fmt.Printf("Exercise removed from routine: %s\n", name)
	return nil
}

func GetExercisesInRoutine(routineID string) ([]structmodels.Exercise, error) {
	query := `SELECT e.id, e.name, e.description 
              FROM Exercises e
              INNER JOIN RoutineExercises re ON e.id = re.exercise_id
              WHERE re.routine_id = ?`

	rows, err := db.Query(query, routineID)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var exercises []structmodels.Exercise
	for rows.Next() {
		var exercise structmodels.Exercise
		if err := rows.Scan(&exercise.ID, &exercise.Name, &exercise.Description); err != nil {
			return nil, fmt.Errorf("failed to scan exercise: %v", err)
		}
		exercises = append(exercises, exercise)
	}

	// Verificar si ocurrió un error durante la iteración
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return exercises, nil
}
