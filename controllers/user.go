package api

import (
	"gin-easy/service/user"
	"gin-easy/utils/context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserRegisterHandler(c *gin.Context) {
	// 请求绑定
	var req user.RegisterReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		context.Error(c, err)
		return
	}
	// 注册
	err = req.Register()
	if err != nil {
		context.Error(c, err)
		return
	}
	context.Success(c, nil)
	return
}

func UserLoginHandler(c *gin.Context) {
	// 请求绑定
	var req user.LoginReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		context.Error(c, err)
		return
	}
	// 登陆
	token, err := req.Login()
	if err != nil {
		context.Error(c, err)
		return
	}
	context.Success(c, token)
	return
}

func UserGetMeHandler(c *gin.Context) {
	// 获取当前用户信息
	uid, err := context.GetUserID(c)
	if err != nil {
		context.Error(c, err)
		return
	}
	id, _ := primitive.ObjectIDFromHex(uid)

	// 请求绑定
	req := user.BasicReq{ID: id}
	// 获取资料
	profile, err := req.GetProfile()
	if err != nil {
		context.Error(c, err)
		return
	}
	context.Success(c, profile)
	return
}

func UserDeleteHandler(c *gin.Context) {
	// 获取当前用户信息
	uid, err := context.GetUserID(c)
	if err != nil {
		context.Error(c, err)
		return
	}
	id, _ := primitive.ObjectIDFromHex(uid)

	// 请求绑定
	req := user.BasicReq{ID: id}
	// 账户注销
	err = req.Delete()
	if err != nil {
		context.Error(c, err)
		return
	}
	context.Success(c, nil)
	return
}
