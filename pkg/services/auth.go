package services

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/qurabay06/todoapp"
	"github.com/qurabay06/todoapp/pkg/repository"
)

const salt = "hjqrhjqw124617ajfhajs"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
} 