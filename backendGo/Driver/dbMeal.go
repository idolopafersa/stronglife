package driver

import (
	structmodels "backendgo/StructModels"
	"fmt"
)

func GetMeal(id string, userID int) (structmodels.Meal, error) {
	var meal structmodels.Meal
	//fmt.Printf("LLEGO AQUI con ID: %s \n", id)
	query := "SELECT id, description, calories, proteins, fats, carbs, name FROM Meals WHERE id = ? AND user_id = ?"

	err := db.QueryRow(query, id).Scan(&meal.ID, &meal.Description, &meal.Calories, &meal.Proteins, &meal.Fats, &meal.Carbs, &meal.Name, userID)
	fmt.Print(err)

	return meal, err

}

func PostMeal(nmeal structmodels.NewMeal, userID int) (int, error) {
	//For now, all meals are private (default value isPublic=0)
	//We need to add a relation with user_id
	query := "INSERT INTO Meals (name, description, calories, proteins, fats, carbs,user_id, isPublic) VALUES (?,? , ?, ?, ?, ?, ?, ?)"
	result, err := db.Exec(query, nmeal.Name, nmeal.Description, nmeal.Calories, nmeal.Proteins, nmeal.Fats, nmeal.Carbs, userID, 0)
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
	query := `UPDATE Meals SET name = ?, description = ?, calories = ?, proteins = ?, fats = ?, carbs = ? WHERE id = ?`

	_, err := db.Exec(query, meal.Name, meal.Description, meal.Calories, meal.Proteins, meal.Fats, meal.Carbs, meal.ID)
	fmt.Print(err)
	return err
}

func GetAllMeals() ([]structmodels.Meal, error) {
	var meals []structmodels.Meal

	// GetAllMeals retrieves all meals from the database which are public (isPublic = 1)
	rows, err := db.Query("SELECT id, name, description, calories, proteins, fats, carbs FROM Meals where isPublic = 1")
	if err != nil {
		return nil, fmt.Errorf("error querying meals: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var meal structmodels.Meal
		if err := rows.Scan(&meal.ID, &meal.Name, &meal.Description, &meal.Calories, &meal.Proteins, &meal.Fats, &meal.Carbs); err != nil {
			return nil, fmt.Errorf("error scanning meal: %v", err)
		}
		meals = append(meals, meal)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return meals, nil
}
