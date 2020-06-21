package api

import (
	"gin-easy/service/user_service"
	"gin-easy/views"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(c *gin.Context) {
	var req user_service.UserBasicReq
	errValid := reqValidator(c, &req)
	if errValid != nil {
		c.JSON(200, views.ErrorResponse(errValid))
		return
	}
	err := req.Register()
	if err != nil {
		c.JSON(200, views.ErrorResponse(err))
		return
	}
	c.JSON(200, views.Response(nil))
	return
}

func UserLoginHandler(c *gin.Context) {
	var req user_service.UserBasicReq
	errValid := reqValidator(c, &req)
	if errValid != nil {
		c.JSON(200, views.ErrorResponse(errValid))
		return
	}
	token, err := req.Login()
	if err != nil {
		c.JSON(200, views.ErrorResponse(err))
		return
	}
	c.JSON(200, views.Response(token))
	return
}

func UserDeleteHandler(c *gin.Context) {
	var req user_service.UserBasicReq
	errValid := reqValidator(c, &req)
	if errValid != nil {
		c.JSON(200, views.ErrorResponse(errValid))
		return
	}
	err := req.Delete()
	if err != nil {
		c.JSON(200, views.ErrorResponse(err))
		return
	}
	c.JSON(200, views.Response(nil))
	return
}
