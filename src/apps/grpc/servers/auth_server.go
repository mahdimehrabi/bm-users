package servers

import (
	"bm-users/src/apps/grpc/proto/auth"
	"bm-users/src/domain/services"
	"bm-users/src/entities"
	"context"
	"errors"
)

type AuthenticationServer struct {
	authService *services.AuthService
}

func (a AuthenticationServer) Login(ctx context.Context, req *auth.LoginReq) (*auth.TokensResponse, error) {
	user := &entities.User{
		Email:    req.Email,
		Password: req.Password,
	}
	jwt, err := a.authService.Login(user)
	if err != nil {
		if errors.Is(err,)
		return nil, err
	}
	tokenResponse := &auth.TokensResponse{
		AccessToken:     jwt.AccessToken,
		RefreshToken:    jwt.RefreshToken,
		AccessTokenExp:  int64(jwt.AccessTokenExp),
		RefreshTokenExp: int64(jwt.RefreshTokenExp),
	}
	return tokenResponse, nil
}

func (a AuthenticationServer) RenewToken(ctx context.Context, req *auth.TokenReq) (*auth.TokensResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthenticationServer) BanUser(ctx context.Context, req *auth.IDReq) (*auth.Empty, error) {
	//TODO implement me
	panic("implement me")
}
