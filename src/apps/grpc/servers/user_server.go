package servers

import (
	"bm-users/src/apps/grpc/proto/user"
	"bm-users/src/domain/services"
	"bm-users/src/entities"
	"context"
)

type UserServer struct {
	userService *services.UserService
}

func NewUserServer(userService *services.UserService) *UserServer {
	return &UserServer{userService: userService}
}

func (u UserServer) GetUser(ctx context.Context, req *user.IDReq) (*user.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServer) Create(ctx context.Context, req *user.UserReq) (*user.UserResponse, error) {
	userr := &entities.User{
		Email:    req.Email,
		Fullname: req.Fullname,
	}
	err := u.userService.CreateUser(userr)
	if err != nil {
		return nil, err
	}
	ur := &user.UserResponse{
		Email:    userr.Email,
		Fullname: userr.Fullname,
	}
	return ur, err
}

func (u UserServer) Update(ctx context.Context, req *user.UserReq) (*user.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServer) ListUsers(ctx context.Context, empty *user.Empty) (*user.ListUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServer) DeleteUser(ctx context.Context, req *user.UserReq) (*user.ListUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
