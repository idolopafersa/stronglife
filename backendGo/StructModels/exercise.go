package structmodels

type NewExercise struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Exercise struct {
	ID int `json:"id"`
	NewExercise
}
