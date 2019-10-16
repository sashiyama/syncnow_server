package repository

import (
	"database/sql"
	. "github.com/sashiyama/syncnow_server/model"
	"golang.org/x/crypto/bcrypt"
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

func (ur *UserRepository) FindByUser(u User) (string, error) {
	var userId string
	var passwordDigest string

	query := `SELECT users.id, user_credentials.password_digest FROM users
                  INNER JOIN user_credentials ON users.id = user_credentials.user_id
                  WHERE user_credentials.email = $1;`

	err := ur.DB.QueryRow(query, u.Email).Scan(&userId, &passwordDigest)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordDigest), []byte(u.Password))
	if err != nil {
		return "", err
	}

	return userId, err
}
