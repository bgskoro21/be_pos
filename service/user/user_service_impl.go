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
	"time"

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

func generateAndStoreRefreshToken(userId uint, ip string, ua string, repo repositoryRefreshToken.RefreshTokenRepository) (string, error){
	refreshToken, err := helper.GenerateJWTRefreshToken(userId);
	if err != nil {
		return "", err
	}

	refreshTokenRequest := &domain.RefreshToken{
		Token: refreshToken,
		UserID: userId,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
		CreatedAt: time.Now(),
		UserAgent: ua,
		IPAddress: ip,
	}

	token, err := repo.Create(refreshTokenRequest);

	if err != nil {
		return "", nil
	}

	return token.Token, nil
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

func (service *UserServiceImpl) Login(request dto.LoginRequest) (map[string]string, error){
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

	accessToken, err := helper.GenerateJWT(user.ID)
	helper.PanicIfError(err)

	refreshToken, err := generateAndStoreRefreshToken(user.ID, request.IPAddress, request.UserAgent, service.refreshTokenRepository);

	helper.PanicIfError(err)

	tokens := map[string]string{
		"accessToken": accessToken,
		"refreshToken": refreshToken,
	}

	return tokens, nil
}

func (service *UserServiceImpl) Refresh(request dto.RefreshTokenRequest) (map[string]string, error){
	if err := helper.ValidateStruct(request); err != nil{
		return nil, err
	}

	data, err := service.refreshTokenRepository.FindByToken(&domain.RefreshToken{
		Token: request.Token,
		IPAddress: request.IPAddress,
		UserAgent: request.UserAgent,
	})

	if err != nil {
		return nil, exception.NewNotFoundError("Refresh token not found or invalid")
	}

	if time.Now().After(data.ExpiresAt){
		return nil, exception.NewUnAuthorizedError("Refresh token expired")
	}

	accessToken, err := helper.GenerateJWT(data.UserID)

	if err != nil {
		return nil, err
	}

	token, err := generateAndStoreRefreshToken(data.UserID, request.IPAddress, request.UserAgent, service.refreshTokenRepository)

	if err != nil {
		return nil, err
	}

	err = service.refreshTokenRepository.DeleteByToken(data.Token);

	if err != nil {
		return nil, err
	}

	tokens := map[string]string{
		"accessToken": accessToken,
		"refreshToken": token,
	}

	return tokens, nil

}

func (service *UserServiceImpl) FindById(userId uint) (*domain.User, error){
	logger.Log.Info(fmt.Sprintf("Looking user for ID: %v", userId))
	user, err := service.userRepository.FindById(userId)

	if err != nil{
		helper.PanicIfError(exception.NewNotFoundError("User not found!"))
	}

	return user, nil;
}