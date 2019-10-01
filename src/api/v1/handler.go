package v1

import (
	"github.com/sashiyama/syncnow_server/services"
)

type Handler struct {
	UserService services.UserService
}
