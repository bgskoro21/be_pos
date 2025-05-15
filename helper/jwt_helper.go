package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId uint) (string, error){
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

func GenerateJWTRefreshToken(userId uint) (string, error){
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("REFRESH_JWT_SECRET")
	return token.SignedString([]byte(secret))
}