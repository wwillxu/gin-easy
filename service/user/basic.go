package user

import (
	"gin-easy/models"
	"gin-easy/views"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type BasicReq struct {
	ID primitive.ObjectID
}

func (service *BasicReq) GetProfile() (interface{}, int) {
	filter := bson.M{"status": models.Normal, "_id": service.ID}
	res, err := models.UserFindOne(filter)
	if err != nil {
		if err.Error() == models.NotExist {
			return nil, views.ErrCliUserNotExist
		}
		log.Println(err)
		return nil, views.ErrorServer
	}
	return gin.H{
		"username":  res.Username,
		"telephone": res.Telephone,
		"email":     res.Email,
	}, views.Success
}

func (service *BasicReq) Delete() int {
	filter := bson.M{"_id": service.ID}
	update := bson.M{"status": models.Deleted}
	_, err := models.UserUpdateOneOfSet(filter, update)
	if err != nil {
		log.Println(err)
		return views.ErrorServer
	}
	return views.Success
}
