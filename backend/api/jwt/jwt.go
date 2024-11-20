package jwt

import (
	"github.com/SwanHtetAungPhyo/api/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)



func GenerateJWT(userID string, TokenType string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"type": TokenType,
		"userID": userID,
		"exp":    expirationTime.Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(config.JwtSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
