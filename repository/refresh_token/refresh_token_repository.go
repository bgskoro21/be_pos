package repository

import "bgskoro21/be-pos/model/domain"

type RefreshTokenRepository interface{
	Create(token *domain.RefreshToken) (*domain.RefreshToken, error)
}

