package service

import (
	"gineasy/models"
	"gineasy/pkg/e"
	"gineasy/pkg/util"
)

type RegisterMsg struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	InviteCode string `json:"invite_code" binding:"required"`
}

func (service *RegisterMsg) valid() int {
	//check username
	if err := models.DB.Model(&models.User{}).Where("username = ?", service.Username).First(&models.User{}).Error; err == nil {
		return e.ErrorUsernameExist
	}
	return e.Success
}

func (service *RegisterMsg) Register() int {
	//valid check
	if err := service.valid(); err != e.Success {
		return err
	}
	//register
	service.Password = util.EncodeMd5(service.Password)
	user := models.User{
		Username:     service.Username,
		Password:     service.Password,
		Nickname:     "anonymity",
		Introduction: "Hi!",
		Avatar:       "http://img.mp.sohu.com/upload/20170619/e1d65856a6f94518a7fba553766015ad_th.png",
	}
	if err := models.DB.Create(&user).Error; err != nil {
		return e.ErrorRegister
	}
	return e.Success
}
