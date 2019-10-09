package repository

import (
	"database/sql"
	. "github.com/sashiyama/syncnow_server/model"
	"time"
)

type RefreshTokenRepository struct {
	DB *sql.DB
}

type RefreshTokenCreateParam struct {
	Tx            *sql.Tx
	AccessTokenId string
	At            time.Time
}

func (rtr *RefreshTokenRepository) Create(p RefreshTokenCreateParam) (RefreshToken, error) {
	const intervalUntilExpiration = 720 * time.Hour
	var err error
	refreshToken := RefreshToken{}
	query := "INSERT INTO user_refresh_tokens(id, user_access_token_id, token, expires_at) VALUES(DEFAULT, $1, DEFAULT, $2) RETURNING id, token, expires_at"

	if p.Tx != nil {
		err = p.Tx.QueryRow(query, p.AccessTokenId, p.At.Add(intervalUntilExpiration)).Scan(&refreshToken.Id, &refreshToken.Token, &refreshToken.ExpiresAt)
	} else {
		err = rtr.DB.QueryRow(query, p.AccessTokenId, p.At.Add(intervalUntilExpiration)).Scan(&refreshToken.Id, &refreshToken.Token, &refreshToken.ExpiresAt)
	}
	return refreshToken, err
}
