package repository

import (
	"database/sql"
	. "github.com/sashiyama/syncnow_server/model"
	"time"
)

type AccessTokenRepository struct {
	DB *sql.DB
}

type AccessTokenParam struct {
	Tx     *sql.Tx
	UserId string
	At     time.Time
}

const accessTokenIntervalUntilExpiration = 24 * time.Hour

func (atr *AccessTokenRepository) Create(p AccessTokenParam) (AccessToken, error) {
	var err error
	accessToken := AccessToken{}
	query := "INSERT INTO user_access_tokens(id, user_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2) RETURNING id, token, expires_at;"

	if p.Tx != nil {
		err = p.Tx.QueryRow(query, p.UserId, p.At.Add(accessTokenIntervalUntilExpiration)).Scan(&accessToken.Id, &accessToken.Token, &accessToken.ExpiresAt)
	} else {
		err = atr.DB.QueryRow(query, p.UserId, p.At.Add(accessTokenIntervalUntilExpiration)).Scan(&accessToken.Id, &accessToken.Token, &accessToken.ExpiresAt)
	}
	return accessToken, err
}

func (atr *AccessTokenRepository) Update(p AccessTokenParam) (AccessToken, error) {
	var err error
	accessToken := AccessToken{}
	query := "UPDATE user_access_tokens SET token = DEFAULT, expires_at = $1 WHERE user_id = $2 RETURNING id, token, expires_at;"

	if p.Tx != nil {
		err = p.Tx.QueryRow(query, p.At.Add(accessTokenIntervalUntilExpiration), p.UserId).Scan(&accessToken.Id, &accessToken.Token, &accessToken.ExpiresAt)
	} else {
		err = atr.DB.QueryRow(query, p.At.Add(accessTokenIntervalUntilExpiration), p.UserId).Scan(&accessToken.Id, &accessToken.Token, &accessToken.ExpiresAt)
	}
	return accessToken, err
}

func (atr *AccessTokenRepository) FindByToken(token string) (AccessToken, error) {
	accessToken := AccessToken{}
	query := "SELECT id, token, expires_at FROM user_access_tokens WHERE token = $1;"
	err := atr.DB.QueryRow(query, token).Scan(&accessToken.Id, &accessToken.Token, &accessToken.ExpiresAt)
	return accessToken, err
}
