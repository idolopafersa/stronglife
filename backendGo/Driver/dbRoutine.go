package driver

import (
	structmodels "backendgo/StructModels"
	"fmt"
)

func PostRoutine(newRoutine structmodels.NewRoutine) (int, error) {
	query := "INSERT INTO Routines (name, description) VALUES (?, ?, ?)"
	result, err := db.Exec(query, newRoutine.Name, newRoutine.Description)

	if err != nil {
		fmt.Println("Error creating new routine:", err)
		return 0, err
	} else {
		id, _ := result.LastInsertId()
		return int(id), nil
	}

}

func PutRoutine(routine structmodels.Routine) error {
	query := `UPDATE Routines SET name = ?, description = ?,  WHERE id = ? AND user_id = ?`
	_, err := db.Exec(query, routine.Name, routine.Description, routine.ID, routine.userID)
	if err != nil {
		fmt.Println("Error updating routine:", err)
		return err
	}
	return nil
}

func GetRoutine(id string) (structmodels.Routine, error) {
	var routine structmodels.Routine
	query := "SELECT id, name, description FROM Routines WHERE id = ? AND user_id = ?"
	err := db.QueryRow(query, id, userID).Scan(&routine.ID, &routine.Name, &routine.Description)
	if err != nil {
		fmt.Println("Error getting routine:", err)
		return routine, err
	}
	return routine, nil

}

func DelRoutine(id string) error {
	var routineID int
	queryGetID := "SELECT id FROM Routines WHERE id = ? AND user_id = ?"
	db.QueryRow(queryGetID, id, userID).Scan(&routineID)

	query := "DELETE FROM RoutineExercises WHERE routine_id = ?;"
	query2 := "DELETE FROM Routines where id = ?"
	_, err := db.Exec(query, routineID)
	fmt.Print(err)
	_, err2 := db.Exec(query2, id)
	fmt.Print(err2)
	return err
}

func GetAlRoutines() ([]structmodels.Routine, error) {
	var routines []structmodels.Routine

	rows, err := db.Query("SELECT * FROM Routines WHERE user_id = ?;")
	db.QueryRow(rows, userID) //userID lo pasa el controlador

	if err != nil {

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var routine structmodels.Routine
		if err := rows.Scan(&routine.ID, &routine.Name, &routine.Description); err != nil {
			return nil, err
		}
		routines = append(routines, routine)
	}

	return routines, nil
}
