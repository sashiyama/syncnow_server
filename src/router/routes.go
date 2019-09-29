package router

import (
	"github.com/labstack/echo"
	"github.com/sashiyama/syncnow_server/api/v1"
)

func Routes(e *echo.Echo) {
	e.GET("/", v1.Root)

	v1_prefix := e.Group("/v1")
	v1_prefix.GET("", v1.Root)

	v1_prefix.POST("/users", v1.CreateUser)
}
