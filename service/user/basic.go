package user

import (
	"gin-easy/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type BasicReq struct {
	ID primitive.ObjectID
}

func (service *BasicReq) GetProfile() (interface{}, error) {
	res, err := models.UserFindOneByID(service.ID)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"username":  res.Username,
		"telephone": res.Telephone,
		"email":     res.Email,
	}, nil
}

func (service *BasicReq) Delete() error {
	err := models.UserUpdateOneSetStatus(service.ID, -1)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
