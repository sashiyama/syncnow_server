package router

import (
	"github.com/labstack/echo"
	"github.com/sashiyama/syncnow_server/api/v1"
)

func Routes(e *echo.Echo, v1Handler *v1.Handler) {
	e.GET("/", v1Handler.Root)

	v1Prefix := e.Group("/v1")
	v1Prefix.GET("", v1Handler.Root)

	v1Prefix.POST("/users", v1Handler.CreateUser)

	v1Prefix.GET("/emails/:email", v1Handler.GetRegisteredEmail)

	v1Prefix.POST("/tokens", v1Handler.CreateAuthToken)
	// v1Prefix.PUT("/tokens", v1Handler.UpdateAuthToken)
}
