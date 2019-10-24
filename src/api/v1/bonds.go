package v1

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) CreateBond(c echo.Context) (err error) {
	userId := h.currentUserId(c)
}
