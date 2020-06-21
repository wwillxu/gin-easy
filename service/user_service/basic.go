package user_service

import (
	"gin-easy/config"
	"gin-easy/models"
	"gin-easy/utils"
	"gin-easy/views"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type UserBasicReq struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func (service *UserBasicReq) Register() error {
	_, err := models.UserFindOne(bson.M{"status": 0, "username": service.Username})
	if err == nil {
		return views.ErrCliUserExist
	}
	_, err = models.UserInsertOne(models.User{
		Username: service.Username,
		Password: service.Password,
	})
	if err != nil {
		log.Println(err)
		return views.ErrorServer
	}
	return nil
}

func (service *UserBasicReq) Login() (string, error) {
	_, err := models.UserFindOne(bson.M{"status": 0, "username": service.Username, "password": service.Password})
	if err != nil {
		return "", views.ErrCliLogin
	}
	token, err := utils.GenerateToken(service.Username, []byte(config.JwtKey), 120*time.Hour)
	if err != nil {
		log.Println(err)
		return "", views.ErrorServer
	}
	return token, nil
}

func (service *UserBasicReq) Delete() error {
	filter := bson.M{"status": 0, "username": service.Username, "password": service.Password}
	update := bson.M{"status": 1}
	_, err := models.UserUpdateOne(filter, bson.M{"$set": update}, nil)
	return err
}
