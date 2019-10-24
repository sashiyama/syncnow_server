package v1

import (
	"github.com/labstack/echo"
	. "github.com/sashiyama/syncnow_server/model"
	"net/http"
	"time"
)

func (h *Handler) Root(c echo.Context) error {
	return c.JSON(http.StatusOK, &Message{Message: "There are SyncNow APIs in here."})
}

func (h *Handler) Home(c echo.Context) error {
	requestToken, err := GetRequestTokenFromHeader(c.Request().Header)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	isAuthorized, err := h.UserService.IsAuthorized(requestToken.Token, time.Now())
	if !isAuthorized {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, &Message{Message: "Home"})
}
