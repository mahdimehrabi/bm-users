package services

import (
	"bm-users/src/entities"
	"github.com/golang-jwt/jwt"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}
func (as AuthService) CreateAccessToken(user entities.User, exp int64, secret string, deviceToken string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userId"] = user.ID
	atClaims["exp"] = exp
	atClaims["deviceToken"] = deviceToken
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (as AuthService) CreateRefreshToken(user entities.User, exp int64, secret string, deviceToken string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userId"] = user.ID
	atClaims["exp"] = exp
	atClaims["deviceToken"] = deviceToken
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
