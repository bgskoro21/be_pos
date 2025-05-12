package repository

import (
	"bgskoro21/be-pos/exception"
	"bgskoro21/be-pos/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &UserRepositoryImpl{db}
}

func (repository *UserRepositoryImpl) Create(user *domain.User) (*domain.User, error){
	if user.Email != nil{
		existingUser := domain.User{}
		
		err := repository.db.Where("email = ?", user.Email).First(&existingUser).Error

		if err == nil{
			panic(exception.NewConflictError("Email already exists"))
		}
	}

	if err := repository.db.Create(user).Error; err != nil{
		panic(err)
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error){
	var user domain.User
	err := repository.db.Where("email = ?", email).First(&user).Error

	return &user, err
}