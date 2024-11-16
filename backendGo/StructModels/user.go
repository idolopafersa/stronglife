package structmodels

type User struct {
	ID           int    `json:"id"`
	StripeID     string `json:"stripe_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Streak       int    `json:"streak"`
	HasWorkedOut int    `json:"has_worked_out"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
