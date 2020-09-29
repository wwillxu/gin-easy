package user

import (
	"gin-easy/config"
	"gin-easy/models"
	"gin-easy/utils"
	"gin-easy/views"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type LoginReq struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func (service *LoginReq) Login() (interface{}, error) {
	// 登陆信息校验
	psw := utils.String2md5(service.Password)
	user, err := models.UserFindOne(bson.M{"status": models.Normal, "username": service.Username, "password": psw})
	if err != nil {
		return nil, views.LoginError
	}
	// 生成token
	token, err := utils.GenerateToken(user.ID, []byte(config.JwtKey))
	if err != nil {
		log.Println(err)
		return nil, views.ServerError
	}
	return gin.H{"token": token}, nil
}
