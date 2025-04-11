package models

type ErrorResponse struct {
	Error string `json:"error" example:"error massage"`
}

type Token struct {
	Token string `json:"token"`
}
