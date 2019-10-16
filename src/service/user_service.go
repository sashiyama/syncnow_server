package service

import (
	"database/sql"
	. "github.com/sashiyama/syncnow_server/model"
	"github.com/sashiyama/syncnow_server/repository"
	"time"
)

type UserRepositoryIF interface {
	Create(p repository.UserCreateParam) (string, error)
}

type UserCredentialRepositoryIF interface {
	ExistsByEmail(email string) (bool, error)
	Create(p repository.UserCredentialCreateParam) error
}

type AccessTokenRepositoryIF interface {
	Create(p repository.AccessTokenCreateParam) (AccessToken, error)
}

type RefreshTokenRepositoryIF interface {
	Create(p repository.RefreshTokenCreateParam) (RefreshToken, error)
}

type TransactionRepositoryIF interface {
	Transaction(txFunc func(*sql.Tx) (interface{}, error)) (interface{}, error)
}

type UserService struct {
	UserRepository           UserRepositoryIF
	UserCredentialRepository UserCredentialRepositoryIF
	AccessTokenRepository    AccessTokenRepositoryIF
	RefreshTokenRepository   RefreshTokenRepositoryIF
	TransactionRepository    TransactionRepositoryIF
}

func (us *UserService) IsRegistered(email string) (bool, error) {
	exists, err := us.UserCredentialRepository.ExistsByEmail(email)
	if err != nil {
		return false, err
	}
	return exists, err
}

func (us *UserService) SignUp(u *SignUpUser, at time.Time) (AuthToken, error) {
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
