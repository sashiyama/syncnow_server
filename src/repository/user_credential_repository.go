package repository

import (
	"database/sql"
	. "github.com/sashiyama/syncnow_server/model"
)

type UserCredentialRepository struct {
	DB *sql.DB
}

type UserCredentialCreateParam struct {
	Tx     *sql.Tx
	User   *SignUpUser
	UserId string
}

func (ucr *UserCredentialRepository) ExistsByEmail(email string) (bool, error) {
	var exists bool
	err := ucr.DB.QueryRow("SELECT EXISTS(SELECT id FROM user_credentials WHERE email = $1);", email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return exists, err
}

func (ucr *UserCredentialRepository) Create(p UserCredentialCreateParam) error {
	var stmt *sql.Stmt
	var err error

	query := "INSERT INTO user_credentials(user_id, email, password_digest) VALUES($1, $2, $3)"

	if p.Tx != nil {
		stmt, err = p.Tx.Prepare(query)
	} else {
		stmt, err = ucr.DB.Prepare(query)
	}

	if err != nil {
		return err
	}

	passwordDigest, err := p.User.PasswordDigest()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(p.UserId, p.User.Email, passwordDigest)
	if err != nil {
		return err
	}

	return err
}
