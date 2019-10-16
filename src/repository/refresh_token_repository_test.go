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

func TestRefreshTokenRepositoryCreate(t *testing.T) {
	now := time.Now()
	d := db.NewPostgres()
	rtr := repository.RefreshTokenRepository{DB: d}

	t.Run("When not transaction", func(t *testing.T) {
		var userId string
		var accessTokenId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)
		d.QueryRow("INSERT INTO user_access_tokens(id, user_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2) RETURNING id;", userId, now.Add(24*time.Hour)).Scan(&accessTokenId)
		p := repository.RefreshTokenParam{AccessTokenId: accessTokenId, At: now}

		refreshToken, err := rtr.Create(p)

		assert.Nil(t, err)
		assert.NotEmpty(t, refreshToken.Token)

		r := RefreshToken{}
		d.QueryRow("SELECT id, token, expires_at FROM user_refresh_tokens;").Scan(&r.Id, &r.Token, &r.ExpiresAt)

		assert.Equal(t, refreshToken, r)

		util.TruncateAllTables()
	})

	t.Run("When transaction and creation is successful", func(t *testing.T) {
		var userId string
		var accessTokenId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)
		d.QueryRow("INSERT INTO user_access_tokens(id, user_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2) RETURNING id;", userId, now.Add(24*time.Hour)).Scan(&accessTokenId)

		tr := repository.TransactionRepository{DB: d}
		refreshToken, err := tr.Transaction(func(tx *sql.Tx) (interface{}, error) {
			p := repository.RefreshTokenParam{AccessTokenId: accessTokenId, At: now, Tx: tx}
			refreshToken, err := rtr.Create(p)
			return refreshToken, err
		})

		assert.Nil(t, err)
		assert.NotEmpty(t, refreshToken.(RefreshToken).Token)

		util.TruncateAllTables()
	})

	t.Run("When transaction and creation fails", func(t *testing.T) {
		var userId string
		var accessTokenId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)
		d.QueryRow("INSERT INTO user_access_tokens(id, user_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2) RETURNING id;", userId, now.Add(24*time.Hour)).Scan(&accessTokenId)

		tr := repository.TransactionRepository{DB: d}
		_, err := tr.Transaction(func(tx *sql.Tx) (interface{}, error) {
			p := repository.RefreshTokenParam{AccessTokenId: accessTokenId, At: now, Tx: tx}
			refreshToken, _ := rtr.Create(p)
			return refreshToken, errors.New("User Refresh Token creation fails")
		})

		var id string
		d.QueryRow("SELECT id FROM user_refresh_tokens;").Scan(&id)

		assert.NotNil(t, err)
		assert.Empty(t, id)

		util.TruncateAllTables()
	})
}

func TestRefreshTokenRepositoryUpdate(t *testing.T) {
	now := time.Now()
	d := db.NewPostgres()
	rtr := repository.RefreshTokenRepository{DB: d}

	t.Run("When not transaction", func(t *testing.T) {
		var userId string
		var accessTokenId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)
		d.QueryRow("INSERT INTO user_access_tokens(id, user_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2) RETURNING id;", userId, now.Add(24*time.Hour)).Scan(&accessTokenId)
		d.QueryRow("INSERT INTO user_refresh_tokens(id, user_access_token_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2)", accessTokenId, now.Add(72*time.Hour))
		p := repository.RefreshTokenParam{AccessTokenId: accessTokenId, At: now}

		refreshToken, err := rtr.Update(p)

		assert.Nil(t, err)
		assert.NotEmpty(t, refreshToken.Token)

		util.TruncateAllTables()
	})

	t.Run("When transaction and creation fails", func(t *testing.T) {
		var userId string
		var accessTokenId string
		d.QueryRow("INSERT INTO users(id) VALUES(DEFAULT) RETURNING id;").Scan(&userId)
		d.QueryRow("INSERT INTO user_access_tokens(id, user_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2);", userId, now.Add(24*time.Hour)).Scan(&accessTokenId)
		d.QueryRow("INSERT INTO user_refresh_tokens(id, user_access_token_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2)", accessTokenId, now.Add(72*time.Hour))

		tr := repository.TransactionRepository{DB: d}
		_, err := tr.Transaction(func(tx *sql.Tx) (interface{}, error) {
			p := repository.RefreshTokenParam{AccessTokenId: accessTokenId, At: time.Now(), Tx: tx}
			refreshToken, _ := rtr.Update(p)
			return refreshToken, errors.New("User Refresh Token creation fails")
		})

		assert.NotNil(t, err)

		util.TruncateAllTables()
	})
}
