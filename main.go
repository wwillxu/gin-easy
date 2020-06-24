package main

import (
	api "gin-easy/controllers"
	"gin-easy/middlewares"
	//"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	// 路由注册
	r := gin.Default()

	// 性能检测
	// pprof.Register(r)

	// 全局跨域中间件
	r.Use(middlewares.Cors())

	// 业务路由
	v1:=r.Group("/api/v1")
	v1.POST("/user/register", api.UserRegisterHandler)
	v1.POST("/user/login", api.UserLoginHandler)

	// jwt中间件
	v1.Use(middlewares.Jwt())

	// 需要鉴权的业务路由
	v1.GET("/user/profile", api.UserGetMeHandler)
	v1.DELETE("/user/account", api.UserDeleteHandler)

	// 监听端口
	r.Run()
}
