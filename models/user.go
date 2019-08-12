package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var ctx context.Context

//user message
type User struct {
	Id primitive.ObjectID `bson:"_id"`

	Username     string `json:"username" `
	Password     string `json:"password" `
	Nickname     string `json:"nickname" `
	Introduction string `json:"introduction"`

	TotalCollCount   uint16 `json:"total_coll_count"`
	TotalFollowCount uint16 `json:"total_follow_count"`
	CreateTime       int64  `json:"create_time" bson:"create_time"`
}

// insert one
func InsertOneUser(user User) (*mongo.InsertOneResult, error) {
	res, err := UserColl.InsertOne(ctx, bson.M{
		"username":           user.Username,
		"password":           user.Password,
		"nickname":           "沙雕本雕",
		"introduction":       "说点什么介绍自己吧！",
		"total_coll_count":   0,
		"total_follow_count": 0,
		"create_time":        time.Now().Unix(),
	})
	return res, err
}

// find one
func FindOneUser(filter bson.M) (User, error) {
	Msg := User{}
	result := UserColl.FindOne(ctx, filter)
	err := result.Decode(&Msg)
	return Msg, err
}

// update one
func UpdateOneUser(filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	res, err := UserColl.UpdateOne(ctx, filter, update)
	return res, err
}
