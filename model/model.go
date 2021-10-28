package model

var userMap map[string]string

func init () {
	userMap = make(map[string]string)
}

func AddUser(username string, password string) (bool, string) {
	_, ok := userMap[username]
	if ok {
		return false, "Already exists."
	}
	userMap[username] = password
	return true, "Successful build."
}

func CheckUser(username string, password string) (bool, string) {
	passwordTrue, ok := userMap[username]
	if ok == false {
		return false, "User not exists."
	}
	if passwordTrue != password {
		return false, "Wrong password."
	}
	return true, "Successfully login."
}
