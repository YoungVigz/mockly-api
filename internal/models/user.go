package models

type User struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserAuthRequest struct {
	Username string
	Email    string
	Password string
}

type UserLoginRequest struct {
	Email    string
	Password string
}

type UserModifyRequest struct {
	Password string
}

type UserChangePassword struct {
	Password    string
	NewPassword string
}

type UserResponse struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
