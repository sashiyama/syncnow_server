package repository

import (
	"database/sql"
)

type UserRepository struct {
	Db *sql.DB
}

func (ur *UserRepository) Create() string {
	return "success!!"
}
