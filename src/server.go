package main

import (
	"github.com/labstack/echo"
	"github.com/sashiyama/syncnow_server/middleware"
	"github.com/sashiyama/syncnow_server/router"
)

func main() {
	e := echo.New()
	middleware.Middlewares(e)
	router.Routes(e)
	e.Logger.Fatal(e.Start(":3000"))
}
