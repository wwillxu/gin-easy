package main

import (
	api "gin-easy/controllers"
	_ "gin-easy/docs"
	"gin-easy/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title gin-easy
// @version 1.0
// @description 带有鉴权的接口采用Bearer Token 加载请求header上
func main() {
	// 路由注册
	r := gin.Default()

	// 全局跨域中间件
	r.Use(middlewares.Cors())

	// API文档
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 业务路由
	r.POST("/user/register",api.UserRegisterHandler)
	r.POST("/user/login",api.UserLoginHandler)

	// jwt中间件
	r.Use(middlewares.Jwt())

	// 需要鉴权的业务路由
	r.POST("/user/delete",api.UserDeleteHandler)

	// 监听端口
	r.Run()
}
