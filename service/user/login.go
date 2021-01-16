package user

import (
	"gin-easy/models"
	"gin-easy/utils"
	"gin-easy/views"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type LoginReq struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (service *LoginReq) Login() (interface{}, int) {
	// 登陆信息校验
	psw := utils.String2md5(service.Password)
	user, err := models.UserFindOne(bson.M{"status": models.Normal, "username": service.Username, "password": psw})
	if err != nil {
		return nil, views.ErrLogin
	}
	// 生成token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		log.Println(err)
		return nil, views.ErrServer
	}
	return gin.H{"token": token}, views.Success
}
