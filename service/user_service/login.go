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

type UserLoginReq struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func (service *UserLoginReq) Login() (string, int) {
	user, err := models.UserFindOne(bson.M{"status": 0, "username": service.Username, "password": service.Password})
	if err != nil {
		return "", views.ErrCliLogin
	}
	token, err := utils.GenerateToken(user.ID, []byte(config.JwtKey), 120*time.Hour)
	if err != nil {
		log.Println(err)
		return "", views.ErrorServer
	}
	return token, views.Success
}
