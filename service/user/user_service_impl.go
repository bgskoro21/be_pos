package service

import (
	"bgskoro21/be-pos/helper"
	"bgskoro21/be-pos/model/domain"
	"bgskoro21/be-pos/model/dto"
	repository "bgskoro21/be-pos/repository/user"
)

type UserServiceImpl struct{
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService{
	return &UserServiceImpl{
		userRepository,
	}
}

func (service *UserServiceImpl) Register(request dto.RegisterUserRequest) (*domain.User, error){
	if err := helper.ValidateStruct(request); err != nil{
		panic(err)
	}

	hashedPassword := helper.HashPassword(request.Password)
	
	registerRequest := domain.User{
		Name: request.Name,
		Email: &request.Email,
		Password: &hashedPassword,
	}

	return service.userRepository.Create(&registerRequest);
}