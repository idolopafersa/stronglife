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

func GetMeal(id string) (structmodels.Meal, error) {
	var meal structmodels.Meal
	fmt.Printf("LLEGO AQUI con ID: %s", id)
	query := "SELECT * FROM Meals WHERE id = ?"

	err := db.QueryRow(query, id).Scan(&meal.ID, &meal.Description, &meal.Calories, &meal.Proteins, &meal.Fats, &meal.Carbs, &meal.PhotoURL, &meal.Name)
	fmt.Print(err)

	return meal, err

}

func PostMeal(nmeal structmodels.NewMeal) (int, error) {
	query := "INSERT INTO Meals (name, description, calories, proteins, fats, carbs, photo_url) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := db.Exec(query, nmeal.Name, nmeal.Description, nmeal.Calories, nmeal.Proteins, nmeal.Fats, nmeal.Carbs, nmeal.PhotoURL)
	if err != nil {
		fmt.Print(err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	return int(id), nil

}

func DelMeal(id string) error {
	query := "DELETE FROM Meals WHERE id = ?"
	_, err := db.Exec(query, id)
	fmt.Print(err)
	return err
}

func UpdateMeal(meal structmodels.Meal) error {
	query := `UPDATE Meals SET name = ?, description = ?, calories = ?, proteins = ?, fats = ?, carbs = ?, photo_url = ? WHERE id = ?`

	_, err := db.Exec(query, meal.Name, meal.Description, meal.Calories, meal.Proteins, meal.Fats, meal.Carbs, meal.PhotoURL, meal.ID)
	fmt.Print(err)
	return err
}

func GetAllMeals() ([]structmodels.Meal, error) {
	var meals []structmodels.Meal

	// Make sure to explicitly select the columns
	rows, err := db.Query("SELECT id, name, description, calories, proteins, fats, carbs, photo_url FROM Meals")
	if err != nil {
		return nil, fmt.Errorf("error querying meals: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var meal structmodels.Meal
		// Ensure the order of the columns matches the struct definition
		if err := rows.Scan(&meal.ID, &meal.Name, &meal.Description, &meal.Calories, &meal.Proteins, &meal.Fats, &meal.Carbs, &meal.PhotoURL); err != nil {
			return nil, fmt.Errorf("error scanning meal: %v", err)
		}
		meals = append(meals, meal)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return meals, nil
}

func UploadMealImage(id string, path string, file multipart.File) error {

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

	query := "UPDATE Meals SET photo_url = ? WHERE id = ?"
	_, erro := db.Exec(query, path, id)
	if erro != nil {
		log.Printf("error saving meal path in db : %s", erro)
		return erro
	}
	return nil
}
