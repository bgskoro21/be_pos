package app

import (
	controller "bgskoro21/be-pos/controller/user"
	repository "bgskoro21/be-pos/repository/user"
	service "bgskoro21/be-pos/service/user"

	"gorm.io/gorm"
)

type Container struct{
	UserController controller.UserController
}

func InitContainer(db *gorm.DB) *Container{
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userCtrl := controller.NewUserController(userSvc)

	return &Container{
		UserController: userCtrl,
	}
}