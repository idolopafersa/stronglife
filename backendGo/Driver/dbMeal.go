package driver

import (
	structmodels "backendgo/StructModels"
	"database/sql"
	"fmt"
	"log"
)

func GetMeal(id string, userID int) (structmodels.Meal, error) {
	var meal structmodels.Meal
	var ownerID int

	query := "SELECT id, description, calories, proteins, fats, carbs, name, user_id FROM Meals WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&meal.ID, &meal.Description, &meal.Calories, &meal.Proteins, &meal.Fats, &meal.Carbs, &meal.Name, &ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return meal, fmt.Errorf("no meal found with the given id")
		}
		return meal, fmt.Errorf("error querying meal: %v", err)
	}

	// Verificar si el usuario es el dueño de la comida
	if ownerID != userID {
		return meal, fmt.Errorf("forbidden: user does not have access to this meal")
	}

	return meal, nil
}

func PostMeal(nmeal structmodels.NewMeal, userID int) (int, error) {
	query := "INSERT INTO Meals (name, description, calories, proteins, fats, carbs, user_id, isPublic) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := db.Exec(query, nmeal.Name, nmeal.Description, nmeal.Calories, nmeal.Proteins, nmeal.Fats, nmeal.Carbs, userID, 0)
	if err != nil {
		log.Printf("Error inserting meal: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert id: %v", err)
		return 0, err
	}

	return int(id), nil
}

func DelMeal(id string, userID int) error {
	// Verificar si el usuario es el dueño de la comida
	var ownerID int
	query := "SELECT user_id FROM Meals WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no meal found with the given id")
		}
		return fmt.Errorf("error querying meal: %v", err)
	}

	if ownerID != userID {
		return fmt.Errorf("forbidden: user does not have access to this meal")
	}

	// Eliminar la comida
	query = "DELETE FROM Meals WHERE id = ? AND user_id = ?"
	_, err = db.Exec(query, id, userID)
	if err != nil {
		log.Printf("Error deleting meal: %v", err)
	}
	return err
}

func UpdateMeal(meal structmodels.Meal, userID int) error {
	// Verificar si el usuario es el dueño de la comida
	var ownerID int
	query := "SELECT user_id FROM Meals WHERE id = ?"
	err := db.QueryRow(query, meal.ID).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no meal found with the given id")
		}
		return fmt.Errorf("error querying meal: %v", err)
	}

	if ownerID != userID {
		return fmt.Errorf("forbidden: user does not have access to this meal")
	}

	// Actualizar la comida
	query = "UPDATE Meals SET name = ?, description = ?, calories = ?, proteins = ?, fats = ?, carbs = ? WHERE id = ? AND user_id = ?"
	_, err = db.Exec(query, meal.Name, meal.Description, meal.Calories, meal.Proteins, meal.Fats, meal.Carbs, meal.ID, userID)
	if err != nil {
		log.Printf("Error updating meal: %v", err)
		return err
	}
	return nil
}

func GetAllMeals(userID int) ([]structmodels.Meal, error) {
	var meals []structmodels.Meal

	rows, err := db.Query("SELECT id, name, description, calories, proteins, fats, carbs FROM Meals where  user_id = ?", userID)
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
