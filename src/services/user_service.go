package services

import (
	"github.com/sashiyama/syncnow_server/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func (us *UserService) SignUp() string {
	return us.UserRepository.Create()
}
