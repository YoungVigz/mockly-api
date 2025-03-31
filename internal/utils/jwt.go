package utils

import (
	"os"
	"time"

	"github.com/YoungVigz/mockly-api/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
