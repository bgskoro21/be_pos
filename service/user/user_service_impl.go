package service

import (
	"bgskoro21/be-pos/exception"
	"bgskoro21/be-pos/helper"
	"bgskoro21/be-pos/model/domain"
	"bgskoro21/be-pos/model/dto"
	"bgskoro21/be-pos/pkg/logger"
	repositoryRefreshToken "bgskoro21/be-pos/repository/refresh_token"
	repository "bgskoro21/be-pos/repository/user"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct{
	userRepository repository.UserRepository;
	refreshTokenRepository repositoryRefreshToken.RefreshTokenRepository
}

func NewUserService(userRepository repository.UserRepository, refreshTokenRepository repositoryRefreshToken.RefreshTokenRepository) UserService{
	return &UserServiceImpl{
		userRepository,
		refreshTokenRepository,
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

func (service *UserServiceImpl) FindById(userId uint) (*domain.User, error){
	logger.Log.Info(fmt.Sprintf("Looking user for ID: %v", userId))
	user, err := service.userRepository.FindById(userId)

	if err != nil{
		helper.PanicIfError(exception.NewNotFoundError("User not found!"))
	}

	return user, nil;
}