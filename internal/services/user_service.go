package services

import (
	"echommerce/internal/models/users_model"

	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) HashPassword(pass string) (hashed string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	return string(bytes), err
}

func (s *UserService) CreateUser(data users_model.User) error {
	hashedPass, err := s.HashPassword(data.Password)
	if err != nil {
		return err
	}
	data.Password = hashedPass
	if err := s.DB.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *UserService) Authenticate(data users_model.UserLogin) (userReturn users_model.User, err error) {
	var user users_model.User

	fmt.Println(data.Email)
	fmt.Println(data.Password)
	if err := s.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {
		return users_model.User{}, errors.New("invalid email or password")
	}

	fmt.Println(user.ID)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return users_model.User{}, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *UserService) GenerateJWTToken(user users_model.User) (string, error) {
	claims := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("your-secret-key"))
}
