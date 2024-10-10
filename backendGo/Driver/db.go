package driver

import (
	structmodels "backendgo/StructModels"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func OpenDB() {

	dbdirection := fmt.Sprintf("%s:%s@tcp(localhost:3306)/FitCalendar", os.Getenv("usuariodb"), os.Getenv("passdb"))
	var err error
	db, err = sql.Open("mysql", dbdirection)

	fmt.Println("connecting to... %s", dbdirection)
	if err != nil {
		pp := fmt.Sprintf("Couldn't open database.... %s", err)
		panic(pp)
	}

	// Ping the database to check if the connection is successful
	if err := db.Ping(); err != nil {
		log.Println(err)
		panic("Not opened")
	}
	fmt.Println("Connected to MariaDB!")

}

func CloseDB() {
	if db != nil {
		db.Close()
	} else {
		log.Println("Couldnt close database")
	}
}

func CreateUser(username, email, password string) int {

	passwordh, err := hashPassword(password)
	if err != nil {
		log.Println("Couldnt hash password: %s", password)
		return 1
	}

	query := "INSERT INTO Users (username, email,password_hash) VALUES (? , ? , ?)"
	resul, err := db.Exec(query, username, email, passwordh)

	if err != nil {
		log.Println(resul, err)
		return 1
	} else {
		return 0
	}

}

func UserExists(username string) bool {
	var user string

	query := "SELECT username FROM Users WHERE username = ?"

	// Query the database for the user and scan the result into the user struct
	err := db.QueryRow(query, username).Scan(&user)

	if err != nil {
		if err == sql.ErrNoRows {

			return false
		}
		log.Println("Error fetching user: %v\n", err)
		return false
	}

	return true
}

func CorrectPassword(username, password string) bool {
	var aldb []byte
	query := ("SELECT password_hash FROM Users WHERE username = ?")
	fmt.Println(username)
	fmt.Println(password)
	db.QueryRow(query, username).Scan(&aldb)

	if err := bcrypt.CompareHashAndPassword(aldb, []byte(password)); err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

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
func PostRoutine(newRoutine structmodels.NewRoutine) (int, error) {
	query := "INSERT INTO Routines (name, description, photo_url) VALUES (?, ?, ?)"
	result, err := db.Exec(query, newRoutine.Name, newRoutine.Description, newRoutine.PhotoURL)

	if err != nil {
		fmt.Println("Error creating new routine:", err)
		return 0, err
	} else {
		id, _ := result.LastInsertId()
		return int(id), nil
	}

}

func PutRoutine(routine structmodels.Routine) error {
	query := `UPDATE Routines SET name = ?, description = ?, photo_url = ? WHERE id = ?`
	_, err := db.Exec(query, routine.Name, routine.Description, routine.PhotoURL, routine.ID)
	if err != nil {
		fmt.Println("Error updating routine:", err)
		return err
	}
	return nil
}

func GetRoutine(id string) (structmodels.Routine, error) {
	var routine structmodels.Routine
	query := "SELECT id, name, description, photo_url FROM Routines WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&routine.ID, &routine.Name, &routine.Description, &routine.PhotoURL)
	if err != nil {
		fmt.Println("Error getting routine:", err)
		return routine, err
	}
	return routine, nil

}

func DelRoutine(id string) error {
	var routineID int
	queryGetID := "SELECT id FROM Routines WHERE id = ?"
	db.QueryRow(queryGetID, id).Scan(&routineID)

	query := "DELETE FROM RoutineExercises WHERE routine_id = ?;"
	query2 := "DELETE FROM Routines where id = ?"
	_, err := db.Exec(query, routineID)
	fmt.Print(err)
	_, err2 := db.Exec(query2, id)
	fmt.Print(err2)
	return err
}

func AddExerciseToRoutine(routineID, exerciseID string) error {
	query := "INSERT INTO RoutineExercises (routine_id, exercise_id) VALUES (?, ?)"
	_, err := db.Exec(query, routineID, exerciseID)
	return err
}

func RemoveExerciseFromRoutine(routineID, exerciseID string) error {
	query := "DELETE FROM RoutineExercises WHERE routine_id = ? AND exercise_id = ?"
	_, err := db.Exec(query, routineID, exerciseID)
	return err
}

func GetExercisesInRoutine(routineID string) ([]structmodels.Exercise, error) {
	query := `SELECT e.* FROM Exercises e
              INNER JOIN RoutineExercises re ON e.id = re.exercise_id
              WHERE re.routine_id = ?`

	rows, err := db.Query(query, routineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []structmodels.Exercise
	for rows.Next() {
		var exercise structmodels.Exercise
		if err := rows.Scan(&exercise.ID, &exercise.Name, &exercise.Sets, &exercise.Repetitions, &exercise.Description, &exercise.PhotoURL); err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}

	return exercises, nil
}

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
	newDay.UserId = userID // Convert string to int
	newDay.Date = date

	return newDay, int(dayID), nil
}

func GetMealsForDay(dayID int) ([]structmodels.Meal, error) {
	query := `
        SELECT m.id, m.name, m.description, m.calories, m.proteins, m.fats, m.carbs, m.photo_url
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
		if err := rows.Scan(&meal.ID, &meal.Name, &meal.Description, &meal.Calories, &meal.Proteins, &meal.Fats, &meal.Carbs, &meal.PhotoURL); err != nil {
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

func GetUserID(name string) (int, error) {
	var id int
	fmt.Print(name)
	query := `Select id FROM Users WHERE username=?`
	err := db.QueryRow(query, name).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			print(err)
			return 0, err
		}
		print(err)
		return 0, err
	}

	return id, nil
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

func GetAlRoutines() ([]structmodels.Routine, error) {
	var routines []structmodels.Routine

	rows, err := db.Query("SELECT * FROM Routines")
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var routine structmodels.Routine
		if err := rows.Scan(&routine.ID, &routine.Name, &routine.Description, &routine.PhotoURL); err != nil {
			return nil, err
		}
		routines = append(routines, routine)
	}

	return routines, nil
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
