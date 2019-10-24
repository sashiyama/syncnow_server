package v1

import (
	"github.com/sashiyama/syncnow_server/service"
)

type Handler struct {
	UserService service.UserService
}
