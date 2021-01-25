package user

import (
	"errors"
	"gin-easy/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type RegisterReq struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
}

func (service *RegisterReq) Register() error {
	// 检查用户名合法性
	_, err := models.UserFindOneByName(service.Username)
	if err == nil {
		return errors.New("user exist")
	}
	// 密码加密
	pwd, _ := bcrypt.GenerateFromPassword([]byte(service.Password), bcrypt.DefaultCost)
	// 用户信息注册
	_, err = models.UserInsertOne(models.User{
		Username:  service.Username,
		Password:  string(pwd),
		Telephone: service.Telephone,
		Email:     service.Email,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
