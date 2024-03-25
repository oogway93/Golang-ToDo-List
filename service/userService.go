package service

import (
	"crypto/sha512"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"todo_list/repository"
	"todo_list/structs"
)

const (
	salt     = "k20tG+_12sdLWQ8fU"
	tokenTTL = 12 * time.Hour
	key      = "6d2L8Ll2gH!u^"
)

type AuthService struct {
	repo repository.User
}

func NewUserRepository(repo repository.User) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user structs.User) error {
	user.Password = HashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, HashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"ExpiresAt": time.Now().Add(tokenTTL).Unix(),
			"IssuedAt":  time.Now().Unix(),
			"Id":        user.Id,
		})
	return token.SignedString([]byte(key))
}

func HashPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
