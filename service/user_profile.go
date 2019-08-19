package service

import (
	"gineasy/models"
	"gineasy/views"
)

func GetProfile(id string) (*views.Profile,error){
	var userMsg models.User
	//err:=models.DB.Model(&models.User{}).Where("id = ?", id).First(&userMsg).Error
	err:=models.DB.First(&userMsg, id).Error
	return views.BuildProfile(userMsg),err
}
