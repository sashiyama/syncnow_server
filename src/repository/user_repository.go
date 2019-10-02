package repository

import (
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func (ur *UserRepository) Create() string {
	return "success!!"
}
