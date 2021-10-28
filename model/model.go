package model

import (
	"WebServer/util"
	"strconv"
)

var passwordMap map[string]string
var sessionMap map[string]string
//var loggedInMap map[string]string

func init () {
	passwordMap = make(map[string]string)
	sessionMap = make(map[string]string)
	//loggedInMap = make(map[string]string)
}

func AddUser(username string, password string) (bool, string) {
	_, ok := passwordMap[username]
	if ok {
		return false, "Already exists."
	}
	passwordMap[username] = password
	sessionMap[username] = ""
	//loggedInMap[username] = "no"
	return true, "Successful build."
}

func LoginUser(username string, password string) (bool, string) {
	passwordTrue, ok := passwordMap[username]
	if ok == false {
		return false, "User not exists."
	}
	if passwordTrue != password {
		return false, "Wrong password."
	}
	sessionMap[username] = strconv.Itoa(util.GetRandom())
	//loggedInMap[username] = "yes"
	return true, "Successfully login."
}

func CheckUser (username string, session string) (bool, string) {
	sessionTrue, ok := sessionMap[username]
	if ok == false {
		return false, "User not exists."
	}
	if sessionTrue != session {
		return false, "Haven't login in."
	}
	return true, "Good user."
}

func GetSessionId (username string) string {
	return sessionMap[username]
}
