package api

import (
	"gin-easy/service/user_service"
	"gin-easy/views"
	"github.com/gin-gonic/gin"
)

// @Summary 用户注册
// @Produce  json
// @Param  req body user_service.UserBasicReq true "用户信息"
// @Success 200 {array} views.basicResponse
// @Router /user/register [post]
func UserRegisterHandler(c *gin.Context) {
	var req user_service.UserBasicReq
	errValid := reqValidator(c, &req)
	if errValid !=nil{
		c.JSON(200,views.ErrorResponse(errValid))
		return
	}
	err := req.Register()
	if err !=nil{
		c.JSON(200,views.ErrorResponse(err))
		return
	}
	c.JSON(200,views.Response(nil))
	return
}

// @Summary 用户登录
// @Produce  json
// @Param  req body user_service.UserBasicReq true "用户信息"
// @Success 200 {array} views.basicResponse
// @Router /user/login [post]
func UserLoginHandler(c *gin.Context) {
	var req user_service.UserBasicReq
	errValid := reqValidator(c, &req)
	if errValid !=nil{
		c.JSON(200,views.ErrorResponse(errValid))
		return
	}
	token, err := req.Login()
	if err !=nil{
		c.JSON(200,views.ErrorResponse(err))
		return
	}
	c.JSON(200,views.Response(token))
	return
}

// @Summary 用户账户删除
// @Produce  json
// @Param  req body user_service.UserBasicReq true "用户信息"
// @Success 200 {array} views.basicResponse
// @Router /user/delete [post]
// @Security BearerToken
func UserDeleteHandler(c *gin.Context) {
	var req user_service.UserBasicReq
	errValid := reqValidator(c, &req)
	if errValid !=nil{
		c.JSON(200,views.ErrorResponse(errValid))
		return
	}
	err := req.Delete()
	if err !=nil{
		c.JSON(200,views.ErrorResponse(err))
		return
	}
	c.JSON(200,views.Response(nil))
	return
}