package util

import (
	"OnlinePhotoAlbum/views"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetId(c *gin.Context) string  {
	id, err := c.Get("id")
	if !err {
		c.JSON(200, views.ErrorResponse("获取登录信息失败"))
		return ""
	}
	stringId := fmt.Sprintf("%v", id)
	return stringId
}
func EncodePsw(psw string) string {
	pswMd5 := md5.New()
	pswMd5.Write([]byte(psw))
	psw = string(pswMd5.Sum(nil))
	return psw
}
