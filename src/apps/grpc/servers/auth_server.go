package servers

import (
	"bm-users/src/apps/grpc/proto/auth"
	"bm-users/src/domain/services"
	"context"
)

type AuthenticationServer struct {
	AuthServer *services.UserService
}

func (a AuthenticationServer) Login(ctx context.Context, req *auth.LoginReq) (*auth.TokensResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthenticationServer) GetCredits(ctx context.Context, req *auth.TokenReq) (*auth.Credits, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthenticationServer) RenewToken(ctx context.Context, req *auth.TokenReq) (*auth.TokensResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthenticationServer) BanUser(ctx context.Context, req *auth.IDReq) (*auth.Empty, error) {
	//TODO implement me
	panic("implement me")
}
