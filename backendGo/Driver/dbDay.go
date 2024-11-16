package driver

import (
	structmodels "backendgo/StructModels"
	"database/sql"
	"fmt"
)

func GetDay(userID int, date string) (structmodels.Day, error) {
	var requestDay structmodels.Day

	query := `SELECT id, user_id, date, routine_id
          FROM Days
          WHERE user_id = ? AND date = ?`

	err := db.QueryRow(query, userID, date).Scan(&requestDay.Id, &requestDay.UserId, &requestDay.Date, &requestDay.RoutineID)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Print(userID)
			requestDay, dayID, err := CreateEmptyDay(userID, date)
			if err != nil {
				return requestDay, err
			}
			requestDay.Id = dayID
		} else {
			return requestDay, err
		}
	}

	// Then retrieve associated meals for that day
	requestDay.Meals, err = GetMealsForDay(requestDay.Id)
	if err != nil {
		return requestDay, err
	}

	return requestDay, nil
}

func CreateEmptyDay(userID int, date string) (structmodels.Day, int, error) {
	var newDay structmodels.Day
	query := "INSERT INTO Days (user_id, date) VALUES (?, ?)"
	fmt.Println(query, userID, date)
	result, err := db.Exec(query, userID, date)
	if err != nil {
		fmt.Println(err)
		return newDay, 0, err
	}

	dayID, err := result.LastInsertId()
	if err != nil {
		return newDay, 0, err
	}

	newDay.Id = int(dayID)
	newDay.UserId = userID // Convert string to intd
	newDay.Date = date

	return newDay, int(dayID), nil
}

func GetMealsForDay(dayID int) ([]structmodels.Meal, error) {
	query := `
        SELECT m.id, m.name, m.description, m.calories, m.proteins, m.fats, m.carbs
        FROM Meals m
        INNER JOIN DayMeals dm ON m.id = dm.meal_id
        WHERE dm.day_id = ?`

	meals := []structmodels.Meal{}

	rows, err := db.Query(query, dayID)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error meals for day %d: %v", dayID, err)
	}
	defer rows.Close()

	for rows.Next() {
		var meal structmodels.Meal
		if err := rows.Scan(&meal.ID, &meal.Name, &meal.Description, &meal.Calories, &meal.Proteins, &meal.Fats, &meal.Carbs); err != nil {
			return nil, fmt.Errorf("error scanning meal: %v", err)
		}
		meals = append(meals, meal)
	}

	return meals, nil
}

func AddMealToDay(date, userID, mealID string) error {
	var dayid int
	queryid := "SELECT id FROM Days WHERE user_id = ? AND date = ? "

	db.QueryRow(queryid, userID, date).Scan(&dayid)

	query := "INSERT INTO DayMeals (day_id, meal_id) VALUES (?, ?)"
	_, err := db.Exec(query, dayid, mealID)
	return err
}

func DelMealToDay(date, userID, mealID string) error {
	var dayid int
	queryid := "SELECT id FROM Days WHERE user_id = ? AND date = ? "

	db.QueryRow(queryid, userID, date).Scan(&dayid)

	query := "DELETE FROM DayMeals WHERE day_id = ? AND meal_id = ?"
	_, err := db.Exec(query, dayid, mealID)
	return err
}

func AddRoutineToDay(date, userID, routineID string) error {
	fmt.Printf("Date being used: %s\n", date)
	var dayid int
	queryid := "SELECT id FROM Days WHERE user_id = ? AND date = ? "

	db.QueryRow(queryid, userID, date).Scan(&dayid)

	query := "Update Days SET routine_id = ? WHERE id = ?"
	_, erre := db.Exec(query, routineID, dayid)
	fmt.Print(erre)
	return erre
}
