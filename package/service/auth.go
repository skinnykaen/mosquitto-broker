package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/skinnykaen/mqtt-broker"
	"github.com/skinnykaen/mqtt-broker/package/repository"
	"os"
	"time"
)

const (
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct{
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user mqtt.User) (error) {
	user.UserData.Password = generatePasswordHash(user.UserData.Password)
	if(s.repo.FindUser(user.UserData.Email)){
		return errors.New("User already exist")
	}
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return  "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims {
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			Issuer: string(time.Now().Unix()),
		},
		user.Id,
	})

	return token.SignedString([]byte(os.Getenv("JWTSIGNINKEY")))
}

func (s *AuthService) ParseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("JWTSIGNINKEY")), nil
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

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}