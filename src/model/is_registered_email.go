package model

type IsRegisteredEmail struct {
	Email        string `validate:"required,email"`
	IsRegistered bool   `validate:"required"`
}
