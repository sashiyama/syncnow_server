package repository

import (
	"database/sql"
	. "github.com/sashiyama/syncnow_server/model"
	"time"
)

type AccessTokenRepository struct {
	DB *sql.DB
}

type AccessTokenCreateParam struct {
	Tx     *sql.Tx
	UserId string
	At     time.Time
}

func (atr *AccessTokenRepository) Create(p AccessTokenCreateParam) (AccessToken, error) {
	const intervalUntilExpiration = 24 * time.Hour
	var err error
	accessToken := AccessToken{}
	query := "INSERT INTO user_access_tokens(id, user_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2) RETURNING id, token, expires_at;"

	if p.Tx != nil {
		err = p.Tx.QueryRow(query, p.UserId, p.At.Add(intervalUntilExpiration)).Scan(&accessToken.Id, &accessToken.Token, &accessToken.ExpiresAt)
	} else {
		err = atr.DB.QueryRow(query, p.UserId, p.At.Add(intervalUntilExpiration)).Scan(&accessToken.Id, &accessToken.Token, &accessToken.ExpiresAt)
	}
	return accessToken, err
}
