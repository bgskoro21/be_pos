package service

import (
	"bgskoro21/be-pos/model/domain"
	"bgskoro21/be-pos/model/dto"
)

type UserService interface{
	Register(registerUserRequest dto.RegisterUserRequest) (*domain.User, error)
	Login(loginRequest dto.LoginRequest) (map[string]string, error)
	Refresh(request dto.RefreshTokenRequest) (map[string]string, error)
	FindById(userId uint) (*domain.User, error)
}