package repository_test

import (
	"database/sql"
	"errors"
	"github.com/sashiyama/syncnow_server/db"
	"github.com/sashiyama/syncnow_server/model"
	"github.com/sashiyama/syncnow_server/repository"
	"github.com/sashiyama/syncnow_server/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserCredentialRepositoryExistsByEmail(t *testing.T) {
	d := db.NewPostgres()
	ucr := repository.UserCredentialRepository{DB: d}

	t.Run("When not exists", func(t *testing.T) {
		exists, err := ucr.ExistsByEmail("test@example.com")

		assert.Equal(t, exists, false)
		assert.Nil(t, err)
	})

	t.Run("When exists", func(t *testing.T) {
		email := "test@example.com"
		var userId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)
		stmt, _ := d.Prepare("INSERT INTO user_credentials(user_id, email, password_digest) VALUES($1, $2, $3)")
		stmt.Exec(userId, email, "P@ssword")

		exists, err := ucr.ExistsByEmail(email)

		assert.Equal(t, exists, true)
		assert.Nil(t, err)

		util.TruncateAllTables()
	})
}

func TestUserCredentialRepositoryCreate(t *testing.T) {
	d := db.NewPostgres()
	ucr := repository.UserCredentialRepository{DB: d}
	signUpUser := &model.User{Email: "test@example.com", Password: "P@ssword"}

	t.Run("When not transaction", func(t *testing.T) {
		var userId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)

		err := ucr.Create(repository.UserCredentialCreateParam{User: signUpUser, UserId: userId})
		assert.Nil(t, err)

		util.TruncateAllTables()
	})

	t.Run("When transaction and creation is successful", func(t *testing.T) {
		var userId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)

		tr := repository.TransactionRepository{DB: d}
		tr.Transaction(func(tx *sql.Tx) (interface{}, error) {
			err := ucr.Create(repository.UserCredentialCreateParam{User: signUpUser, UserId: userId})
			return nil, err
		})

		var id string
		d.QueryRow("SELECT id FROM user_credentials;").Scan(&id)

		assert.NotEmpty(t, id)

		util.TruncateAllTables()
	})

	t.Run("When transaction and creation fails", func(t *testing.T) {
		var userId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)

		tr := repository.TransactionRepository{DB: d}
		tr.Transaction(func(tx *sql.Tx) (interface{}, error) {
			ucr.Create(repository.UserCredentialCreateParam{Tx: tx, User: signUpUser, UserId: userId})
			return nil, errors.New("User Credential creation fails")
		})

		var id string
		d.QueryRow("SELECT id FROM user_credentials;").Scan(&id)

		assert.Empty(t, id)

		util.TruncateAllTables()
	})
}
