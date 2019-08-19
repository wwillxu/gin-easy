package service

import (
	"gineasy/models"
	"gineasy/pkg/e"
	"gineasy/pkg/jwt"
	"gineasy/pkg/util"
)

type LoginMsg struct {
	Username string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
var uid uint

func (service *LoginMsg) userAuth() int {
	// check user
	var userMsg models.User
	if err:=models.DB.Model(&models.User{}).Where("username = ?", service.Username).First(&userMsg).Error;err!=nil{
		return e.ErrorUsernameNotExist
	}
	// check password
	service.Password = util.EncodeMd5(service.Password)
	if service.Password!=userMsg.Password{
		return e.ErrorPassword
	}
	uid=userMsg.ID
	return e.Success
}

func (service *LoginMsg) Login() (string, int) {
	// user auth
	if err := service.userAuth();err!=e.Success{
		return "",err
	}
	// generate token
	token,err:=jwt.GenerateToken(uid)
	if err!=nil{
		return "",e.ErrorTokenInit
	}
	return token,e.Success
}
