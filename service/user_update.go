package service

import (
	"OnlinePhotoAlbum/models"
	"OnlinePhotoAlbum/views"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewUserMsg struct {
	//Username     string `json:"username" `
	Password     string `json:"password" `
	Nickname     string `json:"nickname" `
	Introduction string `json:"introduction"`
}

func (service *NewUserMsg) Update(id string) *views.Response {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	res, err := models.FindOneUser(filter)
	if err != nil {
		return views.ErrorResponse("用户不存在")
	}
	if service.Nickname!=res.Nickname{
		if _,err:=models.UpdateOneUser(filter,bson.M{"nickname":service.Nickname});err!=nil{
			return views.ErrorResponse("修改失败")
		}
	}
	if service.Introduction!=res.Introduction{
		if _,err:=models.UpdateOneUser(filter,bson.M{"introduction":service.Introduction});err!=nil{
			return views.ErrorResponse("修改失败")
		}
	}
	if service.Password!=res.Password&&service.Password!=""{
		if _,err:=models.UpdateOneUser(filter,bson.M{"password":service.Password});err!=nil{
			return views.ErrorResponse("修改失败")
		}
	}
	return nil
}