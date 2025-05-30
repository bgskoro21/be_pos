package dto

type LoginRequest struct{
	Email		string `json:"email" validate:"required,email"`
	Password	string `json:"password" validate:"required,min=6"`
	UserAgent 	string
	IPAddress	string
}