package model

type Person struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}
