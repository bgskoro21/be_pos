package repository

import (
	"bgskoro21/be-pos/model/domain"
)

type UserRepository interface{
	Create(user *domain.User) (*domain.User, error)
}