package model

import (
	"WebServer/util"
	"errors"
)

var (
	ErrorUsed = errors.New("Username has already being used")
	ErrorNotExist = errors.New("User don't exists")
	ErrorWrongPwd = errors.New("Password isn't correct")
	ErrorNotLogin = errors.New("User hasn't login in")
)

type userInfo struct {
	pwdEncode string
}
var userMap map[string]userInfo //username -> userinfo
var sessionIdMap map[string]string //sessionId -> username

func init () {
	userMap = make(map[string]userInfo)
	sessionIdMap = make(map[string]string)
}

//AddUser add a new user into userinfo
func AddUser(username string, password string) (error) {
	_, ok := userMap[username]
	if ok {
		return ErrorUsed
	}

	userMap[username] = userInfo{
		util.Encode(password),
	}
	return nil
}

//LoginUser deal with a request for user login, and return sessionId if success
func LoginUser(username string, password string) (string, error) {
	user, ok := userMap[username]
	if ok == false {
		return "", ErrorNotExist
	}
	if user.pwdEncode != util.Encode(password) {
		return "", ErrorWrongPwd
	}

	sid := util.GetRandom()
	sessionIdMap[sid] = username
	return sid, nil
}

//CheckUser check if a sessionId exist and login in
func CheckUser (username string, sid string) (error) {
	name, ok := sessionIdMap[sid]
	if ok == false || name != username {
		return ErrorNotLogin
	}

	return nil
}