package repository

import (
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

type UserCreateParam struct {
	Tx *sql.Tx
}

func (ur *UserRepository) Create(p UserCreateParam) (string, error) {
	var userId string
	var err error

	query := "INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;"

	if p.Tx != nil {
		err = p.Tx.QueryRow(query).Scan(&userId)
	} else {
		err = ur.DB.QueryRow(query).Scan(&userId)
	}
	if err != nil {
		return "", err
	}
	return userId, err
}
