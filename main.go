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
	v1 := r.Group("/api/v1")
	v1.POST("/register", api.UserRegisterHandler)
	v1.POST("/login", api.UserLoginHandler)

	// jwt中间件
	v1.Use(middlewares.Jwt())

	// 需要鉴权的业务路由
	user := v1.Group("/user")
	user.GET("", api.UserGetMeHandler)
	user.DELETE("", api.UserDeleteHandler)

	// 监听端口
	r.Run()
}
