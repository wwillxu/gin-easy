package user

import (
	"errors"
	"gin-easy/models"
	"gin-easy/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (service *LoginReq) Login() (interface{}, error) {
	// 登陆信息校验
	user, err := models.UserFindOneByName(service.Username)
	if err != nil {
		return nil, err
	}
	// 身份校验
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(service.Password))
	if err != nil {
		return nil, errors.New("wrong password")
	}
	// 生成token
	token, _ := utils.GenerateToken(user.ID)
	return gin.H{"token": token}, nil
}
