package service

import (
	"errors"
	"time"

	"final-project/entity"
	"final-project/repository"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *entity.User) (string, error)
	Login(email, password string) (string, error)
	GetByEmail(email string) (entity.User, error)
	GetByID(id uint) (entity.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

var jwtKey = []byte("secret_key")

func (s *userService) Register(user *entity.User) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.PasswordHash = string(hash)

	err = s.repo.Create(user)
	if err != nil {
		return "", err
	}

	return s.generateToken(user)
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("email tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("password salah")
	}

	return s.generateToken(&user)
}

func (s *userService) generateToken(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(jwtKey)
}

func (s *userService) GetByEmail(email string) (entity.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) GetByID(id uint) (entity.User, error) {
	return s.repo.FindByID(id)
}
