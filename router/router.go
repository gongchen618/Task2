package router

import (
	"WebServer/controller"
	"WebServer/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RouterInit (e *echo.Echo) {
	//根路径 hello world
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//自用一个新建用户的地址 POST /newUser (字段 username 和 password)
	e.POST("/newUser", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		_, result := model.AddUser(username, password)
		return c.String(http.StatusOK, result)
	})

	//第一阶段：接收 POST /session ；第二阶段：接收 POST /login
	account := e.Group("/v1")
	account.POST("/session", controller.Login)
	account.POST("/login", controller.Pass)
}