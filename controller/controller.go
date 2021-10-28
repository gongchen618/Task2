package controller

import (
	"WebServer/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func NewCookie(name string, value string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	return cookie
}
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	ok, result := model.LoginUser (username, password)
	if ok == false {
		return c.JSON(http.StatusOK, result)
	}

	//c.Response().Header().Add("Location", "http://127.0.0.1:2222/v1/login")
	c.SetCookie(NewCookie("dotcom_user", username))
	c.SetCookie(NewCookie("user_session", model.GetSessionId(username)))
	//c.SetCookie(NewCookie("logged_in", "yes"))

	return c.JSON(http.StatusFound, result)
}

func Pass (c echo.Context) error {
	username, err := c.Cookie("dotcom_user")
	session, err := c.Cookie("user_session")
	if err != nil {
		return c.JSON(http.StatusOK, "Strange input.")
	}

	ok, result := model.CheckUser(username.Value, session.Value)
	if ok == false {
		return c.JSON(http.StatusOK, result)
	}

	c.Response().Header().Add("Location", "http://127.0.0.1:2222/")

	return c.JSON(http.StatusFound, result)
}