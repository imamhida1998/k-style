package request

type Register struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   int    `json:"roleId"`
}
