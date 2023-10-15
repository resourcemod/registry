// Package user provide user service methods
package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/resourcemod/registry/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateAccessToken(user *models.User) error {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))
	user.ExpiredAt = time.Now().AddDate(0, 1, 0).String()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":        user.Name,
		"expired_at": user.ExpiredAt,
	})

	s, err := t.SignedString(key)
	if err != nil {
		return err
	}
	user.AccessToken = s
	return nil
}
