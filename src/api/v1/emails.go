package v1

import (
	"github.com/labstack/echo"
	. "github.com/sashiyama/syncnow_server/model"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func (h *Handler) GetRegisteredEmail(c echo.Context) (err error) {
	email := c.Param("email")
	isRegistered, err := h.UserService.IsRegistered(email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	r := &IsRegisteredEmail{Email: email, IsRegistered: isRegistered}

	if err = c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
	}

	if isRegistered {
		return c.JSON(http.StatusOK, r)
	} else {
		return c.JSON(http.StatusNotFound, r)
	}
}
