package views

import "go.mongodb.org/mongo-driver/bson/primitive"

// basic response
type Response struct {
	Success bool
	Error   string
	Data    interface{}
}

// user profile
type Profile struct {
	Id               primitive.ObjectID
	Username         string
	Nickname         string
	Introduction     string
	TotalCollCount   uint16
	TotalFollowCount uint16
	CreateTime       int64
}
