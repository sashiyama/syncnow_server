package models

type SignUpUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"max=100,min=8"`
}
