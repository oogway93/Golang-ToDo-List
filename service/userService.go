package service

import (
	"golang.org/x/crypto/bcrypt"
	"todo_list/repository"
	"todo_list/structs"
)

const salt = "k20tG+_(8fU"

type AuthService struct {
	repo repository.User
}

func NewUserRepository(repo repository.User) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(user structs.UserAdd) error {
	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPassword
	return s.repo.SignUp(user)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return salt + string(bytes), err
}
