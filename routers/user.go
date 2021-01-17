package routers

import (
	api "gin-easy/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.RouterGroup) {
	user := r.Group("/user")
	user.GET("", api.UserGetMeHandler)
	user.DELETE("", api.UserDeleteHandler)
}
