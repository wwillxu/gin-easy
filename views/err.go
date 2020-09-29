package views

import "errors"

var (
	LoginError   = errors.New("[Client Error] wrong user name or password")
	UserExist    = errors.New("[Client Error] user exist")
	UserNotExist = errors.New("[Client Error] user not exist")

	ServerError = errors.New("[Server Error] contact admin")
)
