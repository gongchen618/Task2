package main

import (
	"WebServer/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	account := e.Group("/user")
	account.POST("/login", Login)
	account.POST("/new", New)

	e.Logger.Fatal(e.Start(":2222"))
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	_, result := model.CheckUser (username, password)
	return c.JSON(http.StatusOK, result)
}

func New(c echo.Context) error {
	// User ID from path `users/:id`
	username := c.FormValue("username")
	password := c.FormValue("password")
	_, result := model.AddUser(username, password)
	return c.String(http.StatusOK, result)
}