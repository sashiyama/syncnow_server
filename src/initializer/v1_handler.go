package initializer

import (
	"github.com/sashiyama/syncnow_server/api/v1"
	"github.com/sashiyama/syncnow_server/db"
	"github.com/sashiyama/syncnow_server/repository"
	"github.com/sashiyama/syncnow_server/service"
)

func V1Handler() v1.Handler {
	d := db.NewPostgres()
	ur := &repository.UserRepository{DB: d}
	ucr := &repository.UserCredentialRepository{DB: d}
	atr := &repository.AccessTokenRepository{DB: d}
	rtr := &repository.RefreshTokenRepository{DB: d}
	tr := &repository.TransactionRepository{DB: d}
	us := service.UserService{
		UserRepository:           ur,
		UserCredentialRepository: ucr,
		AccessTokenRepository:    atr,
		RefreshTokenRepository:   rtr,
		TransactionRepository:    tr,
	}
	return v1.Handler{UserService: us}
}
