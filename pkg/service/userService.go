package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/victorvelo/notes/internal/models"
	"github.com/victorvelo/notes/pkg/repository"
)

const (
	salt    = "fdgdggd34566wdkfjfdsfs"
	signKey = "fdfqrkjkvbfgfgdf4353KSFjH"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type UserService struct {
	repo repository.Authorization
}

func NewUserService(repo repository.Authorization) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Add(user models.User) error {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.Add(&user)
}

func (s *UserService) Get(login, password string) (*models.User, error) {
	return s.repo.Get(login, password)
}

func (s *UserService) generatePasswordHash(password string) string {
	hasher := sha1.New()
	bv := []byte(password)
	hasher.Write(bv)

	return fmt.Sprintf("%x", hasher.Sum([]byte(salt)))
}

func (s *UserService) CreateToken(login, password string) (string, error) {
	user, err := s.repo.Get(login, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(8 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signKey))
}

func (s *UserService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signKey), nil
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

// func (s *AuthService) ListAll() ([]user.User, error) {
// 	return s.repo.ListAll()
// }

// func (s *AuthService) Delete(id int) error {
// 	return s.repo.Delete(id)
// }

// func (s *AuthService) Update(id int, u *user.User) error {
// 	return s.repo.Update(id, u)
// }

// func (s *AuthService) generatePasswordHash(password string) string {
// 	hash := sha1.New()
// 	hash.Write([]byte(password))

// 	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
// }

// func (s *AuthService) Get(login, password string) (*user.User, error) {
// 	return s.repo.Get(login, password)
// }

// func (s *AuthService) GenerateToken(login, password string) (string, error) {
// 	user, err := s.repo.Get(login, s.generatePasswordHash(password))
// 	if err != nil {
// 		return "", err
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 		user.Id,
// 	})

// 	return token.SignedString([]byte(signingKey))
// }
