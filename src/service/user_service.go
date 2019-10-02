package service

import (
	"github.com/sashiyama/syncnow_server/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (us *UserService) SignUp() string {
	return us.UserRepository.Create()
}
