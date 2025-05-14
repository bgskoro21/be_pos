package app

import (
	controller "bgskoro21/be-pos/controller/user"
	repositoryToken "bgskoro21/be-pos/repository/refresh_token"
	repository "bgskoro21/be-pos/repository/user"
	service "bgskoro21/be-pos/service/user"

	"gorm.io/gorm"
)

type Container struct{
	UserController controller.UserController
}

func InitContainer(db *gorm.DB) *Container{
	userRepo := repository.NewUserRepository(db)
	refreshTokenRepo := repositoryToken.NewRefreshTokenRepository(db)
	userSvc := service.NewUserService(userRepo, refreshTokenRepo)
	userCtrl := controller.NewUserController(userSvc)

	return &Container{
		UserController: userCtrl,
	}
}