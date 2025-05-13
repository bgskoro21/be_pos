package repository

import (
	"bgskoro21/be-pos/model/domain"
)

type UserRepository interface{
	Create(user *domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindById(id uint) (*domain.User, error)
}