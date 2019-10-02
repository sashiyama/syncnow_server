package service

import (
	"github.com/sashiyama/syncnow_server/repository"
)

type UserCredentialService struct {
	UserCredentialRepository repository.UserCredentialRepository
}

func (ucs *UserCredentialService) IsRegistered(email string) bool {
	return ucs.UserCredentialRepository.ExistsByEmail(email)
}
