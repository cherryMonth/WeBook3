package model

import "time"

type User struct {
	Id         int       `json:"id"`
	Email      string    `json:"email"`
	UserName   string    `json:"user_name"`
	AboutMe    string    `json:"about_me"`
	LastSeen   time.Time `json:"last_seen"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	AvatarHash string    `json:"avatar_hash"`
}

type LocalAuth struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	Password   string    `json:"password"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
