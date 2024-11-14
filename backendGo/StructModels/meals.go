package structmodels

type NewMeal struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Calories    float64 `json:"calories"`
	Proteins    float64 `json:"proteins"`
	Fats        float64 `json:"fats"`
	Carbs       float64 `json:"carbs"`
}
type Meal struct {
	ID int `json:"id"`
	NewMeal
}
