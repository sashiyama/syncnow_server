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
	"time"
)

func TestAccessTokenRepositoryCreate(t *testing.T) {
	d := db.NewPostgres()
	atr := repository.AccessTokenRepository{DB: d}

	t.Run("When not transaction", func(t *testing.T) {
		var userId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)
		p := repository.AccessTokenParam{UserId: userId, At: time.Now()}

		accessToken, err := atr.Create(p)

		assert.Nil(t, err)
		assert.NotEmpty(t, accessToken.Token)

		r := AccessToken{}
		d.QueryRow("SELECT id, token, expires_at FROM user_access_tokens;").Scan(&r.Id, &r.Token, &r.ExpiresAt)

		assert.Equal(t, accessToken, r)

		util.TruncateAllTables()
	})

	t.Run("When transaction and creation is successful", func(t *testing.T) {
		var userId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)

		tr := repository.TransactionRepository{DB: d}
		accessToken, err := tr.Transaction(func(tx *sql.Tx) (interface{}, error) {
			p := repository.AccessTokenParam{UserId: userId, At: time.Now(), Tx: tx}
			accessToken, err := atr.Create(p)
			return accessToken, err
		})

		assert.Nil(t, err)
		assert.NotEmpty(t, accessToken.(AccessToken).Token)

		util.TruncateAllTables()
	})

	t.Run("When transaction and creation fails", func(t *testing.T) {
		var userId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)

		tr := repository.TransactionRepository{DB: d}
		_, err := tr.Transaction(func(tx *sql.Tx) (interface{}, error) {
			p := repository.AccessTokenParam{UserId: userId, At: time.Now(), Tx: tx}
			accessToken, _ := atr.Create(p)
			return accessToken, errors.New("User Access Token creation fails")
		})

		var id string
		d.QueryRow("SELECT id FROM user_access_tokens;").Scan(&id)

		assert.NotNil(t, err)
		assert.Empty(t, id)

		util.TruncateAllTables()
	})
}

func TestAccessTokenRepositoryUpdate(t *testing.T) {
	now := time.Now()
	d := db.NewPostgres()
	atr := repository.AccessTokenRepository{DB: d}

	t.Run("When not transaction", func(t *testing.T) {
		var userId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)
		d.QueryRow("INSERT INTO user_access_tokens(id, user_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2);", userId, now.Add(24*time.Hour))
		p := repository.AccessTokenParam{UserId: userId, At: time.Now()}

		accessToken, err := atr.Update(p)

		assert.Nil(t, err)
		assert.NotEmpty(t, accessToken.Token)

		util.TruncateAllTables()
	})

	t.Run("When transaction and creation fails", func(t *testing.T) {
		var userId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)
		d.QueryRow("INSERT INTO user_access_tokens(id, user_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2);", userId, now.Add(24*time.Hour))

		tr := repository.TransactionRepository{DB: d}
		_, err := tr.Transaction(func(tx *sql.Tx) (interface{}, error) {
			p := repository.AccessTokenParam{UserId: userId, At: time.Now(), Tx: tx}
			accessToken, _ := atr.Update(p)
			return accessToken, errors.New("User Access Token creation fails")
		})

		assert.NotNil(t, err)

		util.TruncateAllTables()
	})
}
