package router

import (
	"github.com/labstack/echo"
	"github.com/sashiyama/syncnow_server/api/v1"
)

func Routes(e *echo.Echo) {
	e.GET("/", v1.Root)
	e.GET("/v1", v1.Root)
}
