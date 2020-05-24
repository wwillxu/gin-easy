package main

import (
	api "gin-easy/controllers"
	"gin-easy/middlewares"
	"gin-easy/views"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/user/register",api.UserRegisterHandler)
	r.POST("/user/login",api.UserLoginHandler)
	r.POST("/user/delete",api.UserDeleteHandler)
	r.Use(middlewares.Jwt())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200,views.Response("pong!"))
	})
	r.Run()
}
