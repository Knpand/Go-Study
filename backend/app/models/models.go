package models

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Response struct{
	Status string `json:"status"`
	Data string `json:"data"`
	Jwt string `json:"jwt"`
}
type Jwt_Authorized_Response struct{
	Status string `json:"status"`
	Name string `json:"name"`
}