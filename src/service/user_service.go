package service

import (
	"database/sql"
	. "github.com/sashiyama/syncnow_server/model"
	"github.com/sashiyama/syncnow_server/repository"
)

type UserService struct {
	UserRepository           repository.UserRepository
	UserCredentialRepository repository.UserCredentialRepository
	TransactionRepository    repository.TransactionRepository
}

func (us *UserService) IsRegistered(email string) (bool, error) {
	exists, err := us.UserCredentialRepository.ExistsByEmail(email)
	if err != nil {
		return false, err
	}
	return exists, err
}

func (us *UserService) SignUp(u *SignUpUser) (bool, error) {
	err := us.TransactionRepository.Transaction(func(tx *sql.Tx) error {
		userId, err := us.UserRepository.Create(repository.UserCreateParam{Tx: tx})
		if err != nil {
			return err
		}
		err = us.UserCredentialRepository.Create(repository.UserCredentialCreateParam{Tx: tx, User: u, UserId: userId})
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return false, err
	}

	return true, err
}
