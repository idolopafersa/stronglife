package driver

import (
	structmodels "backendgo/StructModels"
	"database/sql"
	"fmt"
	"log"
)

func DelRoutine(id string, userID int) error {
	var ownerID int
	queryGetID := "SELECT user_id FROM Routines WHERE id = ?"
	err := db.QueryRow(queryGetID, id).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no routine found with the given id")
		}
		return fmt.Errorf("error querying routine: %v", err)
	}

	if ownerID != userID {
		return fmt.Errorf("forbidden: user does not have access to this routine")
	}

	query := "DELETE FROM RoutineExercises WHERE routine_id = ?;"
	query2 := "DELETE FROM Routines WHERE id = ?"
	_, err = db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting routine exercises: %v", err)
		return err
	}
	_, err = db.Exec(query2, id)
	if err != nil {
		log.Printf("Error deleting routine: %v", err)
		return err
	}
	return nil
}

func GetAlRoutines(userID int) ([]structmodels.Routine, error) {
	var routines []structmodels.Routine

	query := "SELECT id, name, description FROM Routines WHERE user_id = ?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying routines: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var routine structmodels.Routine
		if err := rows.Scan(&routine.ID, &routine.Name, &routine.Description); err != nil {
			return nil, fmt.Errorf("error scanning routine: %v", err)
		}
		routines = append(routines, routine)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return routines, nil
}

func GetRoutine(id, userID int) (structmodels.Routine, error) {
	var routine structmodels.Routine
	var ownerID int

	query := "SELECT id, name, description, user_id FROM Routines WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&routine.ID, &routine.Name, &routine.Description, &ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return routine, fmt.Errorf("no routine found with the given id")
		}
		return routine, fmt.Errorf("error querying routine: %v", err)
	}

	if ownerID != userID {
		return routine, fmt.Errorf("forbidden: user does not have access to this routine")
	}

	return routine, nil
}

func PostRoutine(nroutine structmodels.NewRoutine, userID int) (int, error) {
	query := "INSERT INTO Routines (name, description, user_id) VALUES (?, ?, ?)"
	result, err := db.Exec(query, nroutine.Name, nroutine.Description, userID)
	if err != nil {
		log.Printf("Error inserting routine: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert id: %v", err)
		return 0, err
	}

	return int(id), nil
}

func PutRoutine(routine structmodels.Routine, userID int) error {
	var ownerID int
	query := "SELECT user_id FROM Routines WHERE id = ?"
	err := db.QueryRow(query, routine.ID).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no routine found with the given id")
		}
		return fmt.Errorf("error querying routine: %v", err)
	}

	if ownerID != userID {
		return fmt.Errorf("forbidden: user does not have access to this routine")
	}

	query = "UPDATE Routines SET name = ?, description = ? WHERE id = ? AND user_id = ?"
	_, err = db.Exec(query, routine.Name, routine.Description, routine.ID, userID)
	if err != nil {
		log.Printf("Error updating routine: %v", err)
		return err
	}
	return nil
}
