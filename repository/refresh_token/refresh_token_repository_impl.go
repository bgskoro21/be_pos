package repository

import (
	"bgskoro21/be-pos/model/domain"

	"gorm.io/gorm"
)

type RefreshTokenRepositoryImpl struct{
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository{
	return &RefreshTokenRepositoryImpl{db}
}

func (r *RefreshTokenRepositoryImpl) Create(token *domain.RefreshToken) (*domain.RefreshToken, error){
	return token, r.db.Create(token).Error
}