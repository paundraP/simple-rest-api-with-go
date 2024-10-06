package middleware

import (
	"fmt"
	"os"
	"paundraP/rest-api-with-go/models"
	"time"

	"github.com/golang-jwt/jwt"
)

var your_secret_key = os.Getenv("SECRET_KEY")
var jwtKey = []byte(your_secret_key)

func GenerateJWT(user models.User) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   fmt.Sprintf("%d", user.ID),
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
