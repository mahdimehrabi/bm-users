package services

import (
	"bm-users/src/domain/repositories/user"
	"bm-users/src/entities"
	"bm-users/src/infrastracture/env"
	"bm-users/src/util"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var ErrNoUserFound = errors.New("no user found with entered credentials")

type AuthService struct {
	userRepository user.UserRepository
	env            *env.Env
}

func NewAuthService(userRepository user.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (as AuthService) HashPassword(user *entities.User) error {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(user.Password + user.Email))
	if err != nil {
		return err
	}
	user.Password = string(hasher.Sum([]byte{}))
	return nil
}

func (as AuthService) Login(loginUser *entities.User) (*JWTDTO, error) {
	users, err := as.userRepository.GetByField("email", loginUser.Email)
	if err != nil {
		fmt.Println("error in get field by id:", err)
		return nil, err
	}
	if len(users) < 1 {
		return nil, ErrNoUserFound
	}
	userEnt := users[0]
	err = as.HashPassword(loginUser)
	if err != nil {
		fmt.Println("error in hash password", err)
		return nil, err
	}

	if userEnt.Password != loginUser.Password {
		return nil, err
	}
	expAt := time.Now().Add(10 * time.Minute).Unix()
	at, err := as.CreateAccessToken(*loginUser, expAt,
		as.env.Secret, util.GenerateRandomString(20))
	if err != nil {
		fmt.Println("error in creating access token", err)
		return nil, err
	}

	expRt := time.Now().Add((30 * 24) * time.Hour).Unix()
	rt, err := as.CreateRefreshToken(*loginUser, expRt,
		as.env.Secret, util.GenerateRandomString(20))
	if err != nil {
		fmt.Println("error in creating access token", err)
		return nil, err
	}

	jwtDTO := &JWTDTO{
		at, rt,
		uint64(expAt), uint64(expRt),
	}

	return jwtDTO, nil
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
