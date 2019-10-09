package service

import (
	"database/sql"
	. "github.com/sashiyama/syncnow_server/model"
	"github.com/sashiyama/syncnow_server/repository"
	"time"
)

type UserService struct {
	UserRepository           repository.UserRepository
	UserCredentialRepository repository.UserCredentialRepository
	AccessTokenRepository    repository.AccessTokenRepository
	RefreshTokenRepository   repository.RefreshTokenRepository
	TransactionRepository    repository.TransactionRepository
}

func (us *UserService) IsRegistered(email string) (bool, error) {
	exists, err := us.UserCredentialRepository.ExistsByEmail(email)
	if err != nil {
		return false, err
	}
	return exists, err
}

func (us *UserService) SignUp(u *SignUpUser) (AuthToken, error) {
	//TODO 現在時刻は外部入力にしたい
	at := time.Now()

	authToken, err := us.TransactionRepository.Transaction(func(tx *sql.Tx) (interface{}, error) {
		userId, err := us.UserRepository.Create(repository.UserCreateParam{Tx: tx})
		if err != nil {
			return nil, err
		}
		err = us.UserCredentialRepository.Create(repository.UserCredentialCreateParam{Tx: tx, User: u, UserId: userId})
		if err != nil {
			return nil, err
		}
		accessToken, err := us.AccessTokenRepository.Create(repository.AccessTokenCreateParam{Tx: tx, UserId: userId, At: at})
		if err != nil {
			return nil, err
		}
		refreshToken, err := us.RefreshTokenRepository.Create(repository.RefreshTokenCreateParam{Tx: tx, AccessTokenId: accessToken.Id, At: at})
		if err != nil {
			return nil, err
		}

		return AuthToken{AccessToken: accessToken.Token, RefreshToken: refreshToken.Token}, err
	})

	if token, ok := authToken.(AuthToken); ok {
		return token, err
	} else {
		return AuthToken{}, err
	}
}
