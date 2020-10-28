package api

import (
	"fmt"
	"gin-easy/service/user"
	"gin-easy/views"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func UserRegisterHandler(c *gin.Context) {
	// 请求绑定
	var req user.RegisterReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, views.ErrResponse(views.ErrParam))
		return
	}
	// 注册
	code := req.Register()
	if code != views.Success {
		c.JSON(http.StatusOK, views.ErrResponse(code))
		return
	}
	c.JSON(http.StatusOK, views.DataResponse(nil))
	return
}

func UserLoginHandler(c *gin.Context) {
	// 请求绑定
	var req user.LoginReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, views.ErrResponse(views.ErrParam))
		return
	}
	// 登陆
	token, code := req.Login()
	if code != views.Success {
		c.JSON(http.StatusOK, views.ErrResponse(code))
		return
	}
	c.JSON(http.StatusOK, views.DataResponse(token))
	return
}

func UserGetMeHandler(c *gin.Context) {
	// 获取当前用户信息
	value, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusOK, views.ErrResponse(views.ErrNilID))
		return
	}
	// 检查ID合法性
	id, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", value))
	if err != nil {
		c.JSON(http.StatusOK, views.ErrResponse(views.ErrInvalidID))
		return
	}

	// 请求绑定
	req := user.BasicReq{ID: id}
	// 获取资料
	profile, code := req.GetProfile()
	if code != views.Success {
		c.JSON(http.StatusOK, views.ErrResponse(code))
		return
	}
	c.JSON(http.StatusOK, views.DataResponse(profile))
	return
}

func UserDeleteHandler(c *gin.Context) {
	// 获取当前用户信息
	value, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusOK, views.ErrResponse(views.ErrNilID))
		return
	}
	// 检查ID合法性
	id, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", value))
	if err != nil {
		c.JSON(http.StatusOK, views.ErrResponse(views.ErrInvalidID))
		return
	}

	// 请求绑定
	req := user.BasicReq{ID: id}
	// 账户注销
	code := req.Delete()
	if code != views.Success {
		c.JSON(http.StatusOK, views.ErrResponse(code))
		return
	}
	c.JSON(http.StatusOK, views.DataResponse(nil))
	return
}
