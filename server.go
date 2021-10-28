package main

import (
	"WebServer/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.RouterInit(e)
	e.Logger.Fatal(e.Start(":2222"))
}
