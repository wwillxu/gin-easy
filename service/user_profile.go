package service

import (
	"OnlinePhotoAlbum/models"
	"OnlinePhotoAlbum/views"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfile(id string) (models.User, *views.Response ){
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	if res, err := models.FindOneUser(filter); err != nil {
		return res,views.ErrorResponse("用户不存在")
	}else {
		return res,nil
	}
}