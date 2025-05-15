package dto

type RefreshTokenRequest struct{
	Token		string  `json:"token" validate:"required,min=6"`
	UserAgent 	string
	IPAddress	string
}