package repository

import (
	"database/sql"
	. "github.com/sashiyama/syncnow_server/model"
	"time"
)

type RefreshTokenRepository struct {
	DB *sql.DB
}

type RefreshTokenParam struct {
	Tx            *sql.Tx
	AccessTokenId string
	At            time.Time
}

const refreshTokenIntervalUntilExpiration = 720 * time.Hour

func (rtr *RefreshTokenRepository) Create(p RefreshTokenParam) (RefreshToken, error) {
	var err error
	refreshToken := RefreshToken{}
	query := "INSERT INTO user_refresh_tokens(id, user_access_token_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2) RETURNING id, token, expires_at;"

	if p.Tx != nil {
		err = p.Tx.QueryRow(query, p.AccessTokenId, p.At.Add(refreshTokenIntervalUntilExpiration)).Scan(&refreshToken.Id, &refreshToken.Token, &refreshToken.ExpiresAt)
	} else {
		err = rtr.DB.QueryRow(query, p.AccessTokenId, p.At.Add(refreshTokenIntervalUntilExpiration)).Scan(&refreshToken.Id, &refreshToken.Token, &refreshToken.ExpiresAt)
	}
	return refreshToken, err
}

func (rtr *RefreshTokenRepository) Update(p RefreshTokenParam) (RefreshToken, error) {
	var err error
	refreshToken := RefreshToken{}
	query := "UPDATE user_refresh_tokens SET token = DEFAULT, expires_at = $1 WHERE user_access_token_id = $2 RETURNING id, token, expires_at;"

	if p.Tx != nil {
		err = p.Tx.QueryRow(query, p.At.Add(refreshTokenIntervalUntilExpiration), p.AccessTokenId).Scan(&refreshToken.Id, &refreshToken.Token, &refreshToken.ExpiresAt)
	} else {
		err = rtr.DB.QueryRow(query, p.At.Add(refreshTokenIntervalUntilExpiration), p.AccessTokenId).Scan(&refreshToken.Id, &refreshToken.Token, &refreshToken.ExpiresAt)
	}
	return refreshToken, err
}
