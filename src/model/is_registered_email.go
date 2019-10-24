package model

type IsRegisteredEmail struct {
	Email        string `json:"email" validate:"required,email"`
	IsRegistered bool   `json:"is_registered"`
}
