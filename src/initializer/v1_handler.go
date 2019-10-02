package initializer

import (
	"github.com/sashiyama/syncnow_server/api/v1"
	"github.com/sashiyama/syncnow_server/db"
	"github.com/sashiyama/syncnow_server/repository"
	"github.com/sashiyama/syncnow_server/service"
)

func V1Handler() v1.Handler {
	d := db.New()
	ur := repository.UserRepository{Db: d}
	us := service.UserService{UserRepository: ur}
	return v1.Handler{UserService: us}
}
