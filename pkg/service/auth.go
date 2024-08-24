package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/Grishun/todo"
	"github.com/Grishun/todo/pkg/repository"
)

const (
	salt       = "aboba"
	signingKey = "example12345678adaegwarghrhgwr9"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	Userid int `json:"user_id"`
}

type AuthService struct {
	rep repository.Rep
}

func NewAuthService(repo repository.Rep) *AuthService {
	return &AuthService{rep: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.genPassHash(user.Password)
	return s.rep.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	psw := s.genPassHash(password)
	user, err := s.rep.GetUser(username, psw)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Userid: user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) genPassHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	return claims.Userid, nil
}
