package structmodels

type Routine struct {
	ID          *int       `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Exercises   []Exercise `json:exercises`
}
