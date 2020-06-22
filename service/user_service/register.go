package user_service

import (
	"gin-easy/models"
	"gin-easy/views"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type UserRegisterReq struct {
	Username  string `validate:"required"`
	Password  string `validate:"required"`
	Email     string `validate:"required"`
	Telephone string `validate:"required"`
}

func (service *UserRegisterReq) Register() int {
	_, err := models.UserFindOne(bson.M{"status": 0, "username": service.Username})
	if err == nil {
		return views.ErrCliUserExist
	}
	_, err = models.UserInsertOne(models.User{
		Username:  service.Username,
		Password:  service.Password,
		Telephone: service.Telephone,
		Email:     service.Email,
	})
	if err != nil {
		log.Println(err)
		return views.ErrorServer
	}
	return views.Success
}
