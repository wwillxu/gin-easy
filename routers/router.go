package routers

import (
	"OnlinePhotoAlbum/conf"
	api "OnlinePhotoAlbum/controllers/v1"
	"OnlinePhotoAlbum/middleware"
	"OnlinePhotoAlbum/views"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//set run mode
	gin.SetMode(conf.GinMode)

	//cors
	r.Use(middleware.Cors())

	//api
	v1:=r.Group("/api/v1")
	{
		v1.POST("/ping",pingHandler)
		v1.POST("/user", api.RegisterHandler)
		v1.POST("/user/login", api.LoginHandler)
		v1.GET("/user/:id",api.UserProfileHandler)
		v1.Use(middleware.GetToken().MiddlewareFunc())
		{
			v1.GET("/me",api.MyProfileHandler)
			v1.POST("/me",api.UpdateProfileHandler)
		}
	}
	return r
}

func pingHandler(c *gin.Context)  {
	c.JSON(200,views.SuccessResponse("pong"))
}