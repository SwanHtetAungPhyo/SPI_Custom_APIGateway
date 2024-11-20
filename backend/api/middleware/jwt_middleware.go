package middleware

import (
	"fmt"
	"github.com/SwanHtetAungPhyo/api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"strings"
)

func JWTMiddleware(ctx *fiber.Ctx) bool {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		log.Println("Authorization header is missing")
		return false
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		log.Println("Invalid authorization format")
		return false
	}

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.JwtSecretKey, nil
	})

	if err != nil {
		log.Println("JWT parsing error:", err)
		return false
	}

	if !token.Valid {
		log.Println("Invalid JWT token")
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Locals("claims", claims)
	}


	return true
}

