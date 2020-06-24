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
		c.JSON(200, views.ErrorResponse(views.ErrCliParam, err))
		return
	}
	// 注册业务
	code := req.Register()
	if code != views.Success {
		c.JSON(200, views.ErrorResponse(code, ""))
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
		c.JSON(200, views.ErrorResponse(views.ErrCliParam, err))
		return
	}
	// 登陆业务
	token, code := req.Login()
	if code != views.Success {
		c.JSON(200, views.ErrorResponse(code, ""))
		return
	}
	c.JSON(200, views.Response(token))
	return
}

func UserGetMeHandler(c *gin.Context) {
	// 获取当前用户信息
	id, err := currentUser(c)
	if err != nil {
		c.JSON(200, views.ErrorResponse(views.ErrCliParam, err))
		return
	}
	// 请求绑定
	req := user_service.UserBasicReq{ID: id}
	// 注销业务
	profile, code := req.GetProfile()
	if code != views.Success {
		c.JSON(200, views.ErrorResponse(code, ""))
		return
	}
	c.JSON(200, views.Response(profile))
	return
}

func UserDeleteHandler(c *gin.Context) {
	// 获取当前用户信息
	id, err := currentUser(c)
	if err != nil {
		c.JSON(200, views.ErrorResponse(views.ErrCliParam, err))
		return
	}
	// 请求绑定
	req := user_service.UserBasicReq{ID: id}
	// 注销业务
	code := req.Delete()
	if code != views.Success {
		c.JSON(200, views.ErrorResponse(code, ""))
		return
	}
	c.JSON(200, views.Response(nil))
	return
}
