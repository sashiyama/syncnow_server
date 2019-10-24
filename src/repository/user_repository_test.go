package repository_test

import (
	"database/sql"
	"errors"
	"github.com/sashiyama/syncnow_server/db"
	. "github.com/sashiyama/syncnow_server/model"
	"github.com/sashiyama/syncnow_server/repository"
	"github.com/sashiyama/syncnow_server/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepositoryCreate(t *testing.T) {
	d := db.NewPostgres()
	ur := repository.UserRepository{DB: d}

	t.Run("When not transaction", func(t *testing.T) {
		userId, err := ur.Create(repository.UserCreateParam{})
		assert.NotNil(t, userId)
		assert.Nil(t, err)
	})

	t.Run("When transaction and creation is successful", func(t *testing.T) {
		tr := repository.TransactionRepository{DB: d}
		tr.Transaction(func(tx *sql.Tx) (interface{}, error) {
			_, err := ur.Create(repository.UserCreateParam{Tx: tx})
			return nil, err
		})

		var userId string
		d.QueryRow("SELECT id FROM users;").Scan(&userId)

		assert.NotEmpty(t, userId)

		util.TruncateAllTables()
	})

	t.Run("When transaction and creation fails", func(t *testing.T) {
		tr := repository.TransactionRepository{DB: d}
		tr.Transaction(func(tx *sql.Tx) (interface{}, error) {
			ur.Create(repository.UserCreateParam{Tx: tx})
			return nil, errors.New("User creation failed")
		})

		var userId string
		d.QueryRow("SELECT id FROM users;").Scan(&userId)

		assert.Empty(t, userId)

		util.TruncateAllTables()
	})
}

func TestUserRepositoryFindByUser(t *testing.T) {
	var userId string
	d := db.NewPostgres()
	ur := repository.UserRepository{DB: d}
	u := User{Email: "test@example.com", Password: "P@ssword!!"}
	passwordDigest, _ := u.PasswordDigest()
	d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)
	d.QueryRow("INSERT INTO user_credentials(id, user_id, email, password_digest) VALUES(DEFAULT, $1, $2, $3) RETURNING id;", userId, u.Email, passwordDigest)

	t.Run("When user exist", func(t *testing.T) {
		uid, err := ur.FindByUser(User{Email: "test@example.com", Password: "P@ssword!!"})

		assert.Equal(t, userId, uid)
		assert.Nil(t, err)
	})

	t.Run("When user do not exist", func(t *testing.T) {
		uid, err := ur.FindByUser(User{Email: "test@example.com", Password: "P@ssword"})

		assert.Empty(t, uid)
		assert.NotNil(t, err)
	})

	util.TruncateAllTables()
}
