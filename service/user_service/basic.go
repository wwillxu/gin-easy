package user_service

import (
	"gin-easy/models"
	"gin-easy/views"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type UserBasicReq struct {
	ID string
}

func (service *UserBasicReq) Delete() int {
	// 检查ID合法性
	id, err := primitive.ObjectIDFromHex(service.ID)
	if err != nil {
		log.Println(service.ID)
		log.Println(err)
		return views.ErrCliInvalidID
	}
	// 软删除账户
	filter := bson.M{"_id": id}
	update := bson.M{"status": 1}
	_, err = models.UserUpdateOneOfSet(filter, update)
	if err != nil {
		log.Println(err)
		return views.ErrorServer
	}
	return views.Success
}
