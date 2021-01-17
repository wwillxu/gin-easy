package main

import (
	"gin-easy/config"
	"gin-easy/middlewares"
	"gin-easy/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 路由注册
	r := gin.Default()

	// 跨域中间件
	r.Use(middlewares.Cors())

	// 基本路由
	index := r.Group(config.Conf.App.Prefix)
	routers.IndexRouters(index)

	// jwt中间件
	index.Use(middlewares.Jwt())

	// 用户路由
	routers.UserRouters(index)

	// 监听端口
	r.Run(config.Conf.App.Port)
}
