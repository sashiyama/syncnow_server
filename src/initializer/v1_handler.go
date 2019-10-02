package initializer

import (
	"github.com/sashiyama/syncnow_server/api/v1"
	"github.com/sashiyama/syncnow_server/db"
	"github.com/sashiyama/syncnow_server/repository"
	"github.com/sashiyama/syncnow_server/service"
)

func V1Handler() v1.Handler {
	d := db.New()
	ur := repository.UserRepository{DB: d}
	ucr := repository.UserCredentialRepository{DB: d}
	us := service.UserService{UserRepository: ur}
	ucs := service.UserCredentialService{UserCredentialRepository: ucr}
	return v1.Handler{UserService: us, UserCredentialService: ucs}
}
