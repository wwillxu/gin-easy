package service

import (
	"OnlinePhotoAlbum/conf"
	"OnlinePhotoAlbum/middleware"
	"OnlinePhotoAlbum/models"
	"OnlinePhotoAlbum/util"
	"OnlinePhotoAlbum/views"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LoginMsg struct {
	Username string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

var IdentityKey = "id"
var userId primitive.ObjectID

func (service *LoginMsg) UserAuth() *views.Response {
	// check username
	filter := bson.M{"username": service.Username}
	if _, err := models.FindOneUser(filter); err != nil {
		return views.ErrorResponse("用户名不存在")
	}
	// check password
	service.Password=util.EncodePsw(service.Password)
	filter = bson.M{"username": service.Username,"password":service.Password}
	if res, err := models.FindOneUser(filter); err != nil {
		return views.ErrorResponse("密码错误")
	}else {
		userId=res.Id
		return nil
	}
}

func (service *LoginMsg) Login() (string,*views.Response){
	if err:=service.UserAuth();err!=nil{
		return "",err
	}
	// create the token
	token := jwt.New(jwt.GetSigningMethod(conf.JwtSignMethod))
	claims := token.Claims.(jwt.MapClaims)
	// write user id into payload
	for key, value := range map[string]interface{}{
		IdentityKey: userId} {
		claims[key] = value
	}
	expire := time.Now().Add(middleware.GetToken().Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = time.Now().Unix()
	tokenString, err := token.SignedString(middleware.GetToken().Key)
	if err != nil {
		return "",views.ErrorResponse("")
	}
	return tokenString,nil
}