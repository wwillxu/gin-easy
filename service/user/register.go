package user

import (
	"gin-easy/models"
	"gin-easy/utils"
	"gin-easy/views"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type RegisterReq struct {
	Username  string `binding:"required"`
	Password  string `binding:"required"`
	Email     string `binding:"required"`
	Telephone string `binding:"required"`
}

func (service *RegisterReq) Register() int {
	// 检查用户名合法性
	_, err := models.UserFindOne(bson.M{"status": models.Normal, "username": service.Username})
	if err == nil {
		return views.ErrUserExist
	}
	// 用户信息注册
	_, err = models.UserInsertOne(models.User{
		Username:  service.Username,
		Password:  utils.String2md5(service.Password),
		Telephone: service.Telephone,
		Email:     service.Email,
	})
	if err != nil {
		log.Println(err)
		return views.ErrServer
	}
	return views.Success
}
