package driver

import (
	structmodels "backendgo/StructModels"
	"database/sql"
	"fmt"
	"log"
)

func GetSet(routineID, exerciseID, setNumber string, userID int) (structmodels.Set, error) {
	var set structmodels.Set
	var ownerID int

	query := "SELECT routine_id, exercise_id, set_number, reps, weight, user_id FROM Sets WHERE routine_id = ? AND exercise_id = ? AND set_number = ?"
	err := db.QueryRow(query, routineID, exerciseID, setNumber).Scan(&set.RoutineID, &set.ExerciseID, &set.SetNumber, &set.Reps, &set.Weight, &ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return set, fmt.Errorf("no set found with the given routine_id, exercise_id, and set_number")
		}
		return set, fmt.Errorf("error querying set: %v", err)
	}

	if ownerID != userID {
		return set, fmt.Errorf("forbidden: user does not have access to this set")
	}

	return set, nil
}

func GetAlSet(routineID, exerciseID string, userID int) ([]structmodels.Set, error) {
	var sets []structmodels.Set

	query := "SELECT routine_id, exercise_id, set_number, reps, weight FROM Sets WHERE routine_id = ? AND exercise_id = ? AND user_id = ?"
	rows, err := db.Query(query, routineID, exerciseID, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying sets: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var set structmodels.Set
		if err := rows.Scan(&set.RoutineID, &set.ExerciseID, &set.SetNumber, &set.Reps, &set.Weight); err != nil {
			return nil, fmt.Errorf("error scanning set: %v", err)
		}
		sets = append(sets, set)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return sets, nil
}

func GetAlSetRoutine(routineID string, userID int) ([]structmodels.Set, error) {
	var sets []structmodels.Set

	query := "SELECT routine_id, exercise_id, set_number, reps, weight FROM Sets WHERE routine_id = ? AND user_id = ?"
	rows, err := db.Query(query, routineID, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying sets: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var set structmodels.Set
		if err := rows.Scan(&set.RoutineID, &set.ExerciseID, &set.SetNumber, &set.Reps, &set.Weight); err != nil {
			return nil, fmt.Errorf("error scanning set: %v", err)
		}
		sets = append(sets, set)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return sets, nil
}

func PutSet(set structmodels.Set, userID int) error {
	var ownerID int
	query := "SELECT user_id FROM Sets WHERE routine_id = ? AND exercise_id = ? AND set_number = ?"
	err := db.QueryRow(query, set.RoutineID, set.ExerciseID, set.SetNumber).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no set found with the given routine_id, exercise_id, and set_number")
		}
		return fmt.Errorf("error querying set: %v", err)
	}

	if ownerID != userID {
		return fmt.Errorf("forbidden: user does not have access to this set")
	}

	query = "UPDATE Sets SET reps = ?, weight = ? WHERE routine_id = ? AND exercise_id = ? AND user_id = ? AND set_number = ?"
	_, err = db.Exec(query, set.Reps, set.Weight, set.RoutineID, set.ExerciseID, userID, set.SetNumber)
	if err != nil {
		log.Printf("Error updating set: %v", err)
		return err
	}
	return nil
}

func PostSet(set structmodels.Set, userID int) error {
	// Verificar si el usuario es el dueño de la rutina y el ejercicio
	var routineOwnerID, exerciseOwnerID int
	query := "SELECT user_id FROM Routines WHERE id = ?"
	err := db.QueryRow(query, set.RoutineID).Scan(&routineOwnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no routine found with the given id")
		}
		return fmt.Errorf("error querying routine: %v", err)
	}

	if routineOwnerID != userID {
		return fmt.Errorf("forbidden: user does not have access to this routine")
	}

	query = "SELECT user_id FROM Exercises WHERE id = ?"
	err = db.QueryRow(query, set.ExerciseID).Scan(&exerciseOwnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no exercise found with the given id")
		}
		return fmt.Errorf("error querying exercise: %v", err)
	}

	if exerciseOwnerID != userID {
		return fmt.Errorf("forbidden: user does not have access to this exercise")
	}

	query = "INSERT INTO Sets (routine_id, exercise_id, set_number, weight, reps, user_id) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, set.RoutineID, set.ExerciseID, set.SetNumber, set.Weight, set.Reps, userID)
	if err != nil {
		log.Printf("Error inserting set: %v", err)
		return err
	}
	return nil
}

func DelSet(routineID, exerciseID, setNumber string, userID int) error {
	var setID int
	var ownerID int
	query := "SELECT id, user_id FROM Sets WHERE routine_id = ? AND exercise_id = ? AND set_number = ?"
	err := db.QueryRow(query, routineID, exerciseID, setNumber).Scan(&setID, &ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no set found with the given routine_id, exercise_id, and set_number")
		}
		return fmt.Errorf("error querying set: %v", err)
	}

	if ownerID != userID {
		return fmt.Errorf("forbidden: user does not have access to this set")
	}

	query = "DELETE FROM Sets WHERE id = ?"
	_, err = db.Exec(query, setID)
	if err != nil {
		log.Printf("Error deleting set: %v", err)
		return err
	}
	return nil
}
