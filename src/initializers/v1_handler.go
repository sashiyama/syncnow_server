package initializers

import (
	"github.com/sashiyama/syncnow_server/api/v1"
	"github.com/sashiyama/syncnow_server/db"
	"github.com/sashiyama/syncnow_server/repositories"
	"github.com/sashiyama/syncnow_server/services"
)

func V1Handler() v1.Handler {
	d := db.New()
	ur := repositories.UserRepository{Db: d}
	us := services.UserService{UserRepository: ur}
	return v1.Handler{UserService: us}
}
