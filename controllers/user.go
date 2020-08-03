package api

import (
	"gin-easy/service/user"
	"gin-easy/views"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(c *gin.Context) {
	// 请求绑定
	var req user.RegisterReq
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
	var req user.LoginReq
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
	req := user.BasicReq{ID: id}
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
	req := user.BasicReq{ID: id}
	// 账户注销
	code := req.Delete()
	if code != views.Success {
		c.JSON(200, views.ErrorResponse(code))
		return
	}
	c.JSON(200, views.Response(nil))
	return
}
