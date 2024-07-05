package request

import "time"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Username  string     `json:"username"`
	Fullname  string     `json:"fullname"`
	Role      string     `json:"role"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
