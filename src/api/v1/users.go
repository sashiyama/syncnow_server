package v1

import (
	"github.com/labstack/echo"
	. "github.com/sashiyama/syncnow_server/models"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func CreateUser(c echo.Context) (err error) {
	sign_up_user := new(SignUpUser)
	if err = c.Bind(sign_up_user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(sign_up_user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
	}

	return c.JSON(http.StatusCreated, sign_up_user)
}
