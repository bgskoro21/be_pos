package service

import (
	"bgskoro21/be-pos/exception"
	"bgskoro21/be-pos/helper"
	"bgskoro21/be-pos/model/domain"
	"bgskoro21/be-pos/model/dto"
	repository "bgskoro21/be-pos/repository/user"

	"golang.org/x/crypto/bcrypt"
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

func (service *UserServiceImpl) Login(request dto.LoginRequest) (string, error){
	if err := helper.ValidateStruct(request); err != nil{
		panic(err)
	}

	user, err := service.userRepository.FindByEmail(request.Email)

	if err != nil{
		panic(exception.NewNotFoundError("Email or password are wrong!"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(request.Password))
	
	if err != nil{
		panic(exception.NewNotFoundError("Email or password are wrong!"))
	}

	return helper.GenerateJWT(user.ID)
}