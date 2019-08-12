package views

import (
	"OnlinePhotoAlbum/models"
)

func SuccessResponse(data interface{}) *Response {
	return &Response{
		Success:true,
		Error: "",
		Data:data,
	}
}

func SuccessProfile(data models.User) *Response {
	return &Response{
		Success: true,
		Error:   "",
		Data:    Profile{
			Id:               data.Id,
			Username:         data.Username,
			Nickname:         data.Nickname,
			Introduction:     data.Introduction,
			TotalCollCount:   data.TotalCollCount,
			TotalFollowCount: data.TotalFollowCount,
			CreateTime:       data.CreateTime,
		},
	}
}