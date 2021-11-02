package router

import (
	"WebServer/controller"
	"github.com/labstack/echo/v4"
)

//routerInitTest build routes for test
func routerInitTest (e *echo.Echo) {
	e.GET("/", controller.HelloWorld)
	e.POST("/newUser", controller.NewUser) //POST /newUser (need username & password)
}

//RouterInit build all routes, including those for test
func RouterInit (e *echo.Echo) {
	routerInitTest(e)

	account := e.Group("/v1")
	account.POST("/session", controller.GetSession)
	account.POST("/login", controller.CheckLogin)
	account.POST("/logout", controller.Logout)
}
