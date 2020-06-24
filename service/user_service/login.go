package user_service

import (
	"gin-easy/config"
	"gin-easy/models"
	"gin-easy/utils"
	"gin-easy/views"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type UserLoginReq struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func (service *UserLoginReq) Login() (interface{}, int) {
	// 登陆信息校验
	psw := utils.String2md5(service.Password)
	user, err := models.UserFindOne(bson.M{"status": 0, "username": service.Username, "password": psw})
	if err != nil {
		return "", views.ErrCliLogin
	}
	// 生成token
	token, err := utils.GenerateToken(user.ID, []byte(config.JwtKey), 120*time.Hour)
	if err != nil {
		log.Println(err)
		return "", views.ErrorServer
	}
	return gin.H{"token":token}, views.Success
}
