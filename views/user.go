package views

import (
	"gineasy/models"
	"time"
)

type Profile struct {
	Id           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	Username     string    `json:"username"`
	Nickname     string    `json:"nickname"`
	Introduction string    `json:"introduction"`
	Avatar       string    `json:"avatar"`
}

func BuildProfile(userMsg models.User) *Profile {
	return &Profile{
		Id:           userMsg.ID,
		CreatedAt:    userMsg.CreatedAt,
		Username:     userMsg.Username,
		Nickname:     userMsg.Nickname,
		Introduction: userMsg.Introduction,
		Avatar:       userMsg.Avatar,
	}
}
