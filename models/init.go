package models


import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserColl *mongo.Collection

//connect database
func MongodbInit(uri,dbname string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	//check connection
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	UserColl = client.Database(dbname).Collection("user")
}