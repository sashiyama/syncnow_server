package v1

import (
	"github.com/labstack/echo"
	"net/http"
)

func Root(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
