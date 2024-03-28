package service

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"todo_list/repository"
	"todo_list/structs"
)

const (
	salt     = "k20tG+_12sdLWQ8fU"
	tokenTTL = 12 * time.Hour
	key      = "6d2L8Ll2gH!u^"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int
}

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix()},
		user.Id})
	return token.SignedString([]byte(key))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(key), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func HashPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
