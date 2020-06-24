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
	r.POST("/user/register", api.UserRegisterHandler)
	r.POST("/user", api.UserLoginHandler)

	// jwt中间件
	r.Use(middlewares.Jwt())

	// 需要鉴权的业务路由
	r.GET("/user", api.UserGetMeHandler)
	r.DELETE("/user", api.UserDeleteHandler)

	// 监听端口
	r.Run()
}
