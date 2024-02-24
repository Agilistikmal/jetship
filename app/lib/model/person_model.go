package model

type Person struct {
	ID        string `json:"id" validate:"required,uuid"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required,e164"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Address   string `json:"address" validate:"required,min=8"`
}
