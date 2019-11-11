package v1

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) Health(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
