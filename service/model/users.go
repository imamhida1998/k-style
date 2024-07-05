package model

import "time"

type User struct {
	Id        string     `json:"id"`
	Username  string     `json:"username"`
	Fullname  string     `json:"fullname"`
	Role      string     `json:"role"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
