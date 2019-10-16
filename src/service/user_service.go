package service

import (
	"database/sql"
	. "github.com/sashiyama/syncnow_server/model"
	"github.com/sashiyama/syncnow_server/repository"
	"time"
)

type UserRepositoryIF interface {
	Create(p repository.UserCreateParam) (string, error)
	FindByUser(u User) (string, error)
}

type UserCredentialRepositoryIF interface {
	ExistsByEmail(email string) (bool, error)
	Create(p repository.UserCredentialCreateParam) error
}

type AccessTokenRepositoryIF interface {
	Create(p repository.AccessTokenParam) (AccessToken, error)
	Update(p repository.AccessTokenParam) (AccessToken, error)
}

type RefreshTokenRepositoryIF interface {
	Create(p repository.RefreshTokenParam) (RefreshToken, error)
	Update(p repository.RefreshTokenParam) (RefreshToken, error)
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

func (us *UserService) SignUp(u *User, at time.Time) (AuthToken, error) {
	authToken, err := us.TransactionRepository.Transaction(func(tx *sql.Tx) (interface{}, error) {
		userId, err := us.UserRepository.Create(repository.UserCreateParam{Tx: tx})
		if err != nil {
			return nil, err
		}
		err = us.UserCredentialRepository.Create(repository.UserCredentialCreateParam{Tx: tx, User: u, UserId: userId})
		if err != nil {
			return nil, err
		}
		accessToken, err := us.AccessTokenRepository.Create(repository.AccessTokenParam{Tx: tx, UserId: userId, At: at})
		if err != nil {
			return nil, err
		}
		refreshToken, err := us.RefreshTokenRepository.Create(repository.RefreshTokenParam{Tx: tx, AccessTokenId: accessToken.Id, At: at})
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

func (us *UserService) SignIn(u *User, at time.Time) (AuthToken, error) {
	userId, err := us.UserRepository.FindByUser(*u)
	if err != nil {
		return AuthToken{}, err
	}

	authToken, err := us.TransactionRepository.Transaction(func(tx *sql.Tx) (interface{}, error) {
		accessToken, err := us.AccessTokenRepository.Update(repository.AccessTokenParam{Tx: tx, UserId: userId, At: at})
		if err != nil {
			return nil, err
		}
		refreshToken, err := us.RefreshTokenRepository.Update(repository.RefreshTokenParam{Tx: tx, AccessTokenId: accessToken.Id, At: at})
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
