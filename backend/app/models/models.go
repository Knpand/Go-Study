package models

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Response struct{
	Status string `json:"status"`
	Data string `json:"data"`
	Jwt string `json:"jwt"`
}