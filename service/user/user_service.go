package service

import (
	"bgskoro21/be-pos/model/domain"
	"bgskoro21/be-pos/model/dto"
)

type UserService interface{
	Register(registerUserRequest dto.RegisterUserRequest) (*domain.User, error)
	Login(loginRequest dto.LoginRequest) (string, error)
}