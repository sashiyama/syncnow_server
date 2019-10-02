package v1

import (
	"github.com/labstack/echo"
	. "github.com/sashiyama/syncnow_server/model"
	"net/http"
)

func (h *Handler) Root(c echo.Context) error {
	return c.JSON(http.StatusOK, &Message{Message: "There are SyncNow APIs in here."})
}
