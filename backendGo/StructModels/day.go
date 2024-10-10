package structmodels

type Day struct {
	Id        int    `json:id`
	UserId    int    `json:user_id`
	Date      string `json:"date"`
	RoutineID *int   `json:"routine_id"`
	Meals     []Meal `json:"meals"`
}
