package controller

import (
	"WebServer/model"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Context struct {
	Err string `json:"error"`
}

var (
	ErrorMissCookie = errors.New ("A cookie is missing")
)

func errorToJSON(err error) Context {
	c := Context{
		err.Error(),
	}
	return c
}

//HelloWorld returns Hello World!
func HelloWorld (c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

//NewUser add a new user
func NewUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	err := model.AddUser(username, password)

	if err != nil{
		return c.JSON(http.StatusBadRequest, errorToJSON(err))
	}
	return c.JSON(http.StatusOK, nil)
}

//GetSession handles a request for login in and get its session if success
func newCookie(name string, value string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	return cookie
}
func GetSession(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	sid, err := model.LoginUser (username, password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorToJSON(err))
	}

	//c.Response().Header().Add("Location", "http://127.0.0.1:2222/v1/login")
	c.SetCookie(newCookie("dotcom_user", username))
	c.SetCookie(newCookie("user_session", sid))
	//c.SetCookie(NewCookie("logged_in", "yes"))
	return c.JSON(http.StatusFound, nil)
}

//CheckLogin checks a user if he has login in
func CheckLogin (c echo.Context) error {
	username, err := c.Cookie("dotcom_user")
	session, err := c.Cookie("user_session")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorToJSON(ErrorMissCookie))
	}

	err = model.CheckUser(username.Value, session.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorToJSON(err))
	}

	c.Response().Header().Add("Location", "http://127.0.0.1:2222/")
	return c.JSON(http.StatusFound, nil)
}

//Logout solves a request for logout
func delCookie(name string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = "delete"
	cookie.Expires = time.Unix(0, 0)
	return cookie
}
func Logout (c echo.Context) error {
	username, err := c.Cookie("dotcom_user")
	session, err := c.Cookie("user_session")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorToJSON(ErrorMissCookie))
	}

	err = model.CheckUser(username.Value, session.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorToJSON(err))
	}

	c.SetCookie(delCookie("dotcom_user"))
	c.SetCookie(delCookie("user_session"))
	c.Response().Header().Add("Location", "http://127.0.0.1:2222/")
	return c.JSON(http.StatusOK, nil)
}