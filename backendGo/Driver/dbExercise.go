package driver

import (
	structmodels "backendgo/StructModels"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func GetExercise(id string) (structmodels.Exercise, error) {
	var resul structmodels.Exercise
	fmt.Print(id)
	query := "SELECT * FROM Exercises WHERE id = ?"

	err := db.QueryRow(query, id).Scan(&resul.ID, &resul.Name, &resul.Sets, &resul.Repetitions, &resul.Description, &resul.PhotoURL)
	return resul, err
}

func PostExercise(nexercise structmodels.NewExercise) (int, error) {

	query := "INSERT INTO Exercises ( name, sets, repetitions, description, photo_url) VALUES ( ?, ?, ?, ?, ?)"
	result, err := db.Exec(query, nexercise.Name, nexercise.Sets, nexercise.Repetitions, nexercise.Description, nexercise.PhotoURL)

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
	query := `UPDATE Exercises SET name = ?, description = ?, sets = ?, repetitions = ?, photo_url = ? WHERE id = ?`
	res, err := db.Exec(query, newexercise.Name, newexercise.Description, newexercise.Sets, newexercise.Repetitions, newexercise.PhotoURL, newexercise.ID)
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
		if err := rows.Scan(&exercise.ID, &exercise.Name, &exercise.Sets, &exercise.Repetitions, &exercise.Description, &exercise.PhotoURL); err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}

	return exercises, nil
}
func UploadExerciseImage(id string, path string, file multipart.File) error {

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	dstFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, file); err != nil {
		return err
	}

	query := "UPDATE Exercises SET photo_url = ? WHERE id = ?"
	_, erro := db.Exec(query, path, id)
	if erro != nil {
		log.Printf("error saving exercise path in db : %s", erro)
		return erro
	}
	return nil
}
