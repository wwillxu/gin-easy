package user

import (
	"gin-easy/models"
	"gin-easy/utils"
	"gin-easy/views"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type RegisterReq struct {
	Username  string `validate:"required"`
	Password  string `validate:"required"`
	Email     string `validate:"required"`
	Telephone string `validate:"required"`
}

func (service *RegisterReq) Register() int {
	// 检查用户名合法性
	_, err := models.UserFindOne(bson.M{"status": 0, "username": service.Username})
	if err == nil {
		return views.ErrCliUserExist
	}
	// 用户信息注册
	psw := utils.String2md5(service.Password)
	_, err = models.UserInsertOne(models.User{
		Username:  service.Username,
		Password:  psw,
		Telephone: service.Telephone,
		Email:     service.Email,
	})
	if err != nil {
		log.Println(err)
		return views.ErrorServer
	}
	return views.Success
}
