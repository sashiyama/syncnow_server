package v1

import (
	"github.com/labstack/echo"
	. "github.com/sashiyama/syncnow_server/model"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func (h *Handler) CreateUser(c echo.Context) (err error) {
	u := new(SignUpUser)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
	}
	_, err = h.UserService.SignUp(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, u)
}
