package tool

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jnates/crud_golang/internal/domain/model"
)

func GenerateToken(user *model.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "tiendita_secret_key" // Fallback for development
	}

	claims := jwt.MapClaims{
		"sub":  user.ID.String(),
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"role": user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
