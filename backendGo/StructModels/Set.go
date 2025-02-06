package structmodels

type Set struct {
	RoutineID  int     `json:"routine_id"`
	ExerciseID int     `json:"exercise_id"`
	SetNumber  int     `json:"set_number"`
	Reps       int     `json:"reps"`
	Weight     float32 `json:"weight"`
}
