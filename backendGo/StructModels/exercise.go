package structmodels

type NewExercise struct {
	Name        string  `json:"name"`
	Sets        *int    `json:"sets"`
	Repetitions *int    `json:"repetitions"`
	Description *string `json:"description"`
	PhotoURL    *string `json:"photo_url"`
}

type Exercise struct {
	ID int `json:"id"`
	NewExercise
}
