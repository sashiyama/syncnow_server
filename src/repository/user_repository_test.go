package repository_test

import (
	"database/sql"
	"errors"
	"github.com/sashiyama/syncnow_server/db"
	"github.com/sashiyama/syncnow_server/repository"
	"github.com/sashiyama/syncnow_server/util"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	util.TruncateAllTables()
	code := m.Run()
	util.TruncateAllTables()
	os.Exit(code)
}

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
		tr.Transaction(func(tx *sql.Tx) error {
			_, err := ur.Create(repository.UserCreateParam{Tx: tx})
			if err != nil {
				return err
			}
			return err
		})

		var userId string
		d.QueryRow("SELECT id FROM users;").Scan(&userId)

		assert.NotNil(t, userId)

		util.TruncateAllTables()
	})

	t.Run("When transaction and creation fails", func(t *testing.T) {
		tr := repository.TransactionRepository{DB: d}
		tr.Transaction(func(tx *sql.Tx) error {
			ur.Create(repository.UserCreateParam{Tx: tx})
			return errors.New("User creation failed")
		})

		var userId string
		d.QueryRow("SELECT id FROM users;").Scan(&userId)

		assert.Equal(t, userId, "")

		util.TruncateAllTables()
	})
}
