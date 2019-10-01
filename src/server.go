package main

import (
	"github.com/labstack/echo"
	"github.com/sashiyama/syncnow_server/initializers"
	"github.com/sashiyama/syncnow_server/middleware"
	"github.com/sashiyama/syncnow_server/router"
	"github.com/sashiyama/syncnow_server/validators"
)

func main() {
	e := echo.New()
	e.Validator = validators.NewValidator()
	middleware.Middlewares(e)
	v1 := initializers.V1Handler()
	router.Routes(e, &v1)
	e.Logger.Fatal(e.Start(":3000"))
}
