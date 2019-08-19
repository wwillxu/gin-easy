package service

import (
	"gineasy/models"
	"gineasy/pkg/e"
	"gineasy/pkg/util"
)

type UpdateMsg struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password"`
	Nickname     string `json:"nickname" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	Avatar       string `json:"avatar" binding:"required"`
}

func (service *UpdateMsg) valid() int {
	//check username
	if err := models.DB.Model(&models.User{}).Where("username = ?", service.Username).First(&models.User{}).Error; err == nil {
		return e.ErrorUsernameExist
	}
	return e.Success
}

func (service *UpdateMsg) Update(id string) int {
	// update
	var userMsg models.User
	if err := models.DB.First(&userMsg, id).Error; err != nil {
		return e.Error
	}
	if userMsg.Username != service.Username {
		//valid check
		if err := service.valid(); err != e.Success {
			return err
		}
	}
	service.Password = util.EncodeMd5(service.Password)
	if err := models.DB.Model(&userMsg).Updates(UpdateMsg{
		Username:     service.Username,
		Password:     service.Password,
		Nickname:     service.Nickname,
		Introduction: service.Introduction,
		Avatar:       service.Avatar,
	}).Error; err != nil {
		return e.ErrorUpdateUser
	}
	return e.Success
}
