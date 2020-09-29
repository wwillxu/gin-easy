package api

import (
	"gin-easy/service/user"
	"gin-easy/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler(c *gin.Context) {
	// 请求绑定
	var req user.RegisterReq
	err := reqValidator(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, views.ErrorResponse(err))
		return
	}
	// 注册
	err = req.Register()
	if err != nil {
		c.JSON(http.StatusOK, views.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, views.Response(nil))
	return
}

func UserLoginHandler(c *gin.Context) {
	// 请求绑定
	var req user.LoginReq
	err := reqValidator(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, views.ErrorResponse(err))
		return
	}
	// 登陆
	token, err := req.Login()
	if err != nil {
		c.JSON(http.StatusOK, views.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, views.Response(token))
	return
}

func UserGetMeHandler(c *gin.Context) {
	// 获取当前用户信息
	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusOK, views.ErrorResponse(err))
		return
	}
	// 请求绑定
	req := user.BasicReq{ID: id}
	// 获取资料
	profile, err := req.GetProfile()
	if err != nil {
		c.JSON(http.StatusOK, views.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, views.Response(profile))
	return
}

func UserDeleteHandler(c *gin.Context) {
	// 获取当前用户信息
	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusOK, views.ErrorResponse(err))
		return
	}
	// 请求绑定
	req := user.BasicReq{ID: id}
	// 账户注销
	err = req.Delete()
	if err != nil {
		c.JSON(http.StatusOK, views.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, views.Response(nil))
	return
}
