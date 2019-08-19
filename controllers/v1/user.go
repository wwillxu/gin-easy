package v1

import (
	"fmt"
	"gineasy/pkg/bind"
	"gineasy/pkg/e"
	"gineasy/service"
	"gineasy/views"
	"github.com/gin-gonic/gin"
)

// user register
func RegisterHandler(c *gin.Context) {
	var newUser service.RegisterMsg
	// bind check
	if err := bind.BindCheck(c, &newUser); err != e.Success {
		c.JSON(200, views.ErrRes(err))
		return
	}
	// register service
	if err := newUser.Register(); err != e.Success {
		c.JSON(200, views.ErrRes(err))
		return
	}
	c.JSON(200, views.SuccessRes("register success"))
}

// user login
func LoginHandler(c *gin.Context) {
	var loginUser service.LoginMsg
	// bind check
	if err := bind.BindCheck(c, &loginUser); err != e.Success {
		c.JSON(200, views.ErrRes(err))
		return
	}
	// login service
	token, err := loginUser.Login()
	if err != e.Success {
		c.JSON(200, views.ErrRes(err))
		return
	}
	c.JSON(200, views.SuccessRes(token))
}

// get user's profile
func UserProfileHandler(c *gin.Context) {
	// bind :id
	var id string
	id = c.Param("id")
	// get profile
	profile,err:=service.GetProfile(id)
	if err!=nil{
		c.JSON(200,views.ErrRes(e.ErrorUsernameNotExist))
		return
	}
	c.JSON(200,views.SuccessRes(profile))

}

// get my profile
func MyProfileHandler(c *gin.Context) {
	// get uid
	var uid string
	if id, err := c.Get("id");!err {
		c.JSON(200, views.ErrRes(e.ErrorGetUID))
		return
	}else {
		uid = fmt.Sprintf("%v", id)
	}
	// get profile
	profile,err:=service.GetProfile(uid)
	if err!=nil{
		c.JSON(200,views.ErrRes(e.Error))
		return
	}
	c.JSON(200,views.SuccessRes(profile))
}

// update user profile
func UpdateProfileHandler(c *gin.Context) {
	var updateUser service.UpdateMsg
	// bind check
	if err := bind.BindCheck(c, &updateUser); err != e.Success {
		c.JSON(200, views.ErrRes(err))
		return
	}
	// get uid
	var uid string
	if id, err := c.Get("id");!err {
		c.JSON(200, views.ErrRes(e.ErrorGetUID))
		return
	}else {
		uid = fmt.Sprintf("%v", id)
	}
	// update service
	if err:=updateUser.Update(uid);err!=e.Success{
		c.JSON(200,views.ErrRes(err))
		return
	}
	c.JSON(200,views.SuccessRes("update success"))
}
