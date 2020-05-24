package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)


type User struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt int64 `json:"created_at" bson:"created_at"`
	DeleteAt int64 `json:"delete_at" bson:"delete_at"`
	Status int `json:"status" bson:"status"`

	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func UserInsertOne(msg User) (*mongo.InsertOneResult, error) {
	res, err := UserColl.InsertOne(context.Background(), bson.M{
		"created_at":time.Now().Unix(),
		"delete_at":time.Now().Unix(),
		"status":0,

		"username":   msg.Username,
		"password":   msg.Password,
	})
	return res, err
}

func UserFindOne(filter bson.M) (User, error) {
	var msg User
	err := UserColl.FindOne(context.Background(), filter).Decode(&msg)
	return msg, err
}

func UserFindMany(filter bson.M, options *options.FindOptions) ([]User, error) {
	ctx := context.Background()
	cursor, err := UserColl.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var res []User
	for cursor.Next(ctx) {
		var temp User
		if err := cursor.Decode(&temp); err != nil {
			return nil, err
		}
		res = append(res, temp)
	}
	return res, nil
}

func UserUpdateOne(filter bson.M, update bson.M, options *options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := UserColl.UpdateOne(context.Background(), filter, update, options)
	return res, err
}

func UserUpdateMany(filter bson.M, update bson.M, options *options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := UserColl.UpdateMany(context.Background(), filter, update, options)
	return res, err
}

func UserDeleteOne(filter bson.M) (*mongo.DeleteResult, error) {
	res, err := UserColl.DeleteOne(context.Background(), filter)
	return res, err
}

func UserDeleteMany(filter bson.M) (*mongo.DeleteResult, error) {
	res, err := UserColl.DeleteMany(context.Background(), filter)
	return res, err
}