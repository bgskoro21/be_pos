package middleware

import (
	"bgskoro21/be-pos/helper"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() fiber.Handler{
	return func(c *fiber.Ctx) error{
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer "){
			return helper.SendResponse(c, fiber.StatusUnauthorized, nil, "Unauthorized")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error){
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid{
			return helper.SendResponse(c, fiber.StatusUnauthorized, nil, "Invalid or expired token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok{
			return helper.SendResponse(c, fiber.StatusUnauthorized, nil, "Invalid token claims")
		}

		c.Locals("user_id", claims["user_id"])

		return c.Next()
	}
}