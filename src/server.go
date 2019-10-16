package main

import (
	"github.com/labstack/echo"
	"github.com/sashiyama/syncnow_server/initializer"
	"github.com/sashiyama/syncnow_server/middleware"
	"github.com/sashiyama/syncnow_server/router"
	"github.com/sashiyama/syncnow_server/validator"
)

func main() {
	e := echo.New()
	e.Validator = validator.NewValidator()
	middleware.Middlewares(e)
	v1 := initializer.V1Handler()
	router.Routes(e, &v1)
	e.Logger.Fatal(e.Start(":80"))
}
