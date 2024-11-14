package util
import (

	"github.com/golang-jwt/jwt/v4"
	"time"
)


func GenerateJWT(secretKey string, userID string) (string, error) {

	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
