package routers

import (
	api "gin-easy/controllers"
	"github.com/gin-gonic/gin"
)

func IndexRouters(r *gin.RouterGroup) {
	r.POST("/register", api.UserRegisterHandler)
	r.POST("/login", api.UserLoginHandler)
}
