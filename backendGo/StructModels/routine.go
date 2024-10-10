package structmodels

type NewRoutine struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	PhotoURL    *string `json:"photo_url"`
}

type Routine struct {
	ID int `json:"id"`
	NewRoutine
}
