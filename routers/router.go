package routers

import (
	"gineasy/conf"
	api "gineasy/controllers/v1"
	"gineasy/middleware"
	"gineasy/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//set run mode
	gin.SetMode(conf.AppRunMode)

	//cors
	r.Use(middleware.Cors())

	//api
	v1:=r.Group("/api/v1")
	{
		v1.POST("/ping",pingHandler)
		v1.POST("/user", api.RegisterHandler)
		v1.POST("/user/login", api.LoginHandler)
		v1.GET("/user/:id",api.UserProfileHandler)
		v1.Use(middleware.Jwt())
		{
			v1.POST("/hello",pingHandler)
			v1.GET("/me",api.MyProfileHandler)
			v1.POST("/me",api.UpdateProfileHandler)
		}
	}
	return r
}

func pingHandler(c *gin.Context)  {
	if uid,err:=c.Get(jwt.PayloadMsg);err{
		c.JSON(200,gin.H{"uid":uid})
		return
	}
	c.JSON(200,gin.H{"message":"pong"})
}