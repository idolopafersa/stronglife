package driver

import (
	structmodels "backendgo/StructModels"
	"database/sql"
	"fmt"
	"log"
)

func GetExercise(id, userID int) (structmodels.Exercise, error) {
	var resul structmodels.Exercise
	var ownerID int

	query := "SELECT id, name, description, user_id FROM Exercises WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&resul.ID, &resul.Name, &resul.Description, &ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return resul, fmt.Errorf("no exercise found with the given id")
		}
		return resul, fmt.Errorf("error querying exercise: %v", err)
	}

	// Verificar si el usuario es el dueño del ejercicio
	if ownerID != userID {
		return resul, fmt.Errorf("forbidden: user does not have access to this exercise")
	}

	return resul, nil
}

func PostExercise(nexercise structmodels.NewExercise, userID int) (int, error) {
	query := "INSERT INTO Exercises (name, description, user_id) VALUES (?, ?, ?)"
	result, err := db.Exec(query, nexercise.Name, nexercise.Description, userID)
	if err != nil {
		log.Printf("Error inserting exercise: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert id: %v", err)
		return 0, err
	}
	return int(id), nil
}

func DelExercise(id, userID int) error {
	// Verificar si el usuario es el dueño del ejercicio
	var ownerID int
	query := "SELECT user_id FROM Exercises WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no exercise found with the given id")
		}
		return fmt.Errorf("error querying exercise: %v", err)
	}

	if ownerID != userID {
		return fmt.Errorf("forbidden: user does not have access to this exercise")
	}

	// Eliminar el ejercicio
	query = "DELETE FROM Exercises WHERE id = ? AND user_id = ?"
	_, err = db.Exec(query, id, userID)
	if err != nil {
		log.Printf("Error deleting exercise: %v", err)
	}
	return err
}

func PutExercise(newexercise structmodels.Exercise, userID int) error {
	// Verificar si el usuario es el dueño del ejercicio
	var ownerID int
	query := "SELECT user_id FROM Exercises WHERE id = ?"
	err := db.QueryRow(query, newexercise.ID).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no exercise found with the given id")
		}
		return fmt.Errorf("error querying exercise: %v", err)
	}

	if ownerID != userID {
		return fmt.Errorf("forbidden: user does not have access to this exercise") //Asi se crean errores personalizados
	}

	// Actualizar el ejercicio
	query = "UPDATE Exercises SET name = ?, description = ? WHERE id = ? AND user_id = ?"
	res, err := db.Exec(query, newexercise.Name, newexercise.Description, newexercise.ID, userID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}
	log.Printf("Rows affected: %d", rowsAffected)
	if rowsAffected == 0 {
		return fmt.Errorf("no exercise found with id: %d", newexercise.ID)
	}
	return nil
}

func GetAlExercises(userID int) ([]structmodels.Exercise, error) {
	var exercises []structmodels.Exercise
	query := fmt.Sprintf("SELECT * FROM Exercises WHERE user_id = %d", userID)

	rows, err := db.Query(query)
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
