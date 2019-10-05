package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Middlewares(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
}
