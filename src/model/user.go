package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"max=100,min=8"`
}

func (u *User) PasswordDigest() (string, error) {
	digest, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(digest), err
}
