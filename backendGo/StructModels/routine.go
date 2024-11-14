package structmodels

type NewRoutine struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Routine struct {
	ID int `json:"id"`
	NewRoutine
}
