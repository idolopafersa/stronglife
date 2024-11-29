package driver

import (
	structmodels "backendgo/StructModels"
	"fmt"
)

func GetExercise(id string) (structmodels.Exercise, error) {
	var resul structmodels.Exercise
	fmt.Print(id)
	query := "SELECT * FROM Exercises WHERE id = ?"

	err := db.QueryRow(query, id).Scan(&resul.ID, &resul.Name, &resul.Description)
	return resul, err
}

func PostExercise(nexercise structmodels.NewExercise, userID int) (int, error) {

	query := "INSERT INTO Exercises ( name, description,user_id) VALUES ( ?, ?, ?)"
	result, err := db.Exec(query, nexercise.Name, nexercise.Description, userID)

	if err != nil {
		fmt.Print(err)
		return 0, err
	} else {
		id, _ := result.LastInsertId()
		return int(id), nil
	}
}

func DelExercise(id string) error {

	query := "DELETE FROM Exercises where id = ?"
	_, err := db.Exec(query, id)
	fmt.Print(err)
	return err
}

func PutExercise(newexercise structmodels.Exercise) error {
	query := `UPDATE Exercises SET name = ?, description = ? WHERE id = ?`
	res, err := db.Exec(query, newexercise.Name, newexercise.Description, newexercise.ID)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Rows affected: %d\n", rowsAffected)
	if rowsAffected == 0 {
		return fmt.Errorf("No exercise found with id: %d", newexercise.ID)
	}
	return nil
}

func GetAlExercises() ([]structmodels.Exercise, error) {
	var exercises []structmodels.Exercise

	rows, err := db.Query("SELECT * FROM Exercises")
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var exercise structmodels.Exercise
		if err := rows.Scan(&exercise.ID, &exercise.Name); err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}

	return exercises, nil
}
