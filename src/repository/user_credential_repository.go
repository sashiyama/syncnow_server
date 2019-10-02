package repository

import (
	"database/sql"
	"log"
)

type UserCredentialRepository struct {
	DB *sql.DB
}

func (ucr *UserCredentialRepository) ExistsByEmail(email string) bool {
	var id string
	err := ucr.DB.QueryRow(`SELECT id FROM user_credentials WHERE email IN (?) LIMIT 1`, email).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return true
}
