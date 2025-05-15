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

func (r *RefreshTokenRepositoryImpl) FindByToken(request *domain.RefreshToken) (*domain.RefreshToken, error){
	var refreshToken domain.RefreshToken;

	err := r.db.Where("token = ?", request.Token).Where("user_agent = ?", request.UserAgent).Where("ip_address = ?", request.IPAddress).First(&refreshToken).Error;

	return &refreshToken, err
}

func (r *RefreshTokenRepositoryImpl) DeleteByToken(token string) error{
	return r.db.Where("token = ?", token).Delete(&domain.RefreshToken{}).Error
}