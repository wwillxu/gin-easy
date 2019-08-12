package v1

import (
	"OnlinePhotoAlbum/service"
	"OnlinePhotoAlbum/views"
	"fmt"
	"github.com/gin-gonic/gin"
)

// user register
func RegisterHandler(c *gin.Context)  {
	var newUser service.SignMsg
	//bind sign massage
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(200, views.ErrorBind(err))
		return
	}
	//register
	if err:=newUser.Register();err!=nil{
		c.JSON(200,err)
		return
	}
	c.JSON(200,views.SuccessResponse("register success"))
}

// user login
func LoginHandler(c *gin.Context)  {
	var loginUser service.LoginMsg
	//bind sign massage
	if err := c.ShouldBind(&loginUser); err != nil {
		c.JSON(200, views.ErrorBind(err))
		return
	}
	//login
	if token,err:=loginUser.Login();err!=nil{
		c.JSON(200,err)
		return
	}else {
		c.JSON(200,views.SuccessResponse(token))
	}
}

// get user's profile
func UserProfileHandler(c *gin.Context)  {
	//bind :id
	var id string
	id = c.Param("id")
	if profile,err:=service.GetProfile(id);err!=nil{
		c.JSON(200,err)
		return
	}else {
		c.JSON(200,views.SuccessProfile(profile))
	}
}

// get my profile
func MyProfileHandler(c *gin.Context)  {
	id, err := c.Get("id")
	if !err {
		c.JSON(200, views.ErrorResponse("获取登录信息失败"))
		return
	}
	stringId := fmt.Sprintf("%v", id)
	if profile,err:=service.GetProfile(stringId);err!=nil{
		c.JSON(200,err)
		return
	}else {
		c.JSON(200,views.SuccessProfile(profile))
	}
}

// update user profile
func UpdateProfileHandler(c *gin.Context)  {
	var newData service.NewUserMsg
	// bind sign massage
	if err := c.ShouldBind(&newData); err != nil {
		c.JSON(200, views.ErrorBind(err))
		return
	}
	// get id
	id, err := c.Get("id")
	if !err {
		c.JSON(200, views.ErrorResponse("获取登录信息失败"))
		return
	}
	stringId := fmt.Sprintf("%v", id)
	// update
	if err:=newData.Update(stringId);err!=nil{
		c.JSON(200,err)
		return
	}
	c.JSON(200,views.SuccessResponse("update success"))
}


