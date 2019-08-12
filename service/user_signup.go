package service

import (
	"OnlinePhotoAlbum/models"
	"OnlinePhotoAlbum/util"
	"OnlinePhotoAlbum/views"
	"go.mongodb.org/mongo-driver/bson"
)

type SignMsg struct {
	Username string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	InviteCode string `json:"invite_code" binding:"required"`
}

func (service *SignMsg) valid() *views.Response {
	// check username
	filter := bson.M{"username": service.Username}
	if _, err := models.FindOneUser(filter); err == nil {
		return views.ErrorResponse("用户名已存在")
	}
	return nil
}

func (service *SignMsg) Register() *views.Response {
	user :=models.User{
		Username:       service.Username,
		Password:       service.Password,
	}
	if err:=service.valid();err!=nil{
		return err
	}
	user.Password=util.EncodePsw(user.Password)
	if _, err := models.InsertOneUser(user); err != nil {
		return views.ErrorResponse("服务器出错导致注册失败")
	}
	return nil
}