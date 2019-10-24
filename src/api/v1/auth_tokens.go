package v1

import (
	"github.com/labstack/echo"
	. "github.com/sashiyama/syncnow_server/model"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"time"
)

func (h *Handler) CreateAuthToken(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
	}
	authToken, err := h.UserService.SignIn(u, time.Now())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, authToken)
}
