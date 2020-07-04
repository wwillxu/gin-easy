package api

import (
	"gin-easy/service/user_service"
	"gin-easy/views"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(c *gin.Context) {
	// 请求绑定
	var req user_service.UserRegisterReq
	err := reqValidator(c, &req)
	if err != nil {
		return
	}
	// 注册
	code := req.Register()
	if code != views.Success {
		c.JSON(200, views.ErrorResponse(code))
		return
	}
	c.JSON(200, views.Response(nil))
	return
}

func UserLoginHandler(c *gin.Context) {
	// 请求绑定
	var req user_service.UserLoginReq
	err := reqValidator(c, &req)
	if err != nil {
		return
	}
	// 登陆
	token, code := req.Login()
	if code != views.Success {
		c.JSON(200, views.ErrorResponse(code))
		return
	}
	c.JSON(200, views.Response(token))
	return
}

func UserGetMeHandler(c *gin.Context) {
	// 获取当前用户信息
	id, err := getUserID(c)
	if err != nil {
		return
	}
	// 请求绑定
	req := user_service.UserBasicReq{ID: id}
	// 获取资料
	profile, code := req.GetProfile()
	if code != views.Success {
		c.JSON(200, views.ErrorResponse(code))
		return
	}
	c.JSON(200, views.Response(profile))
	return
}

func UserDeleteHandler(c *gin.Context) {
	// 获取当前用户信息
	id, err := getUserID(c)
	if err != nil {
		return
	}
	// 请求绑定
	req := user_service.UserBasicReq{ID: id}
	// 账户注销
	code := req.Delete()
	if code != views.Success {
		c.JSON(200, views.ErrorResponse(code))
		return
	}
	c.JSON(200, views.Response(nil))
	return
}
