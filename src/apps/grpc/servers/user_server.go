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
	um, err := u.userService.GetUserByID(req.GetID())
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{
		Email:    um.Email,
		Fullname: um.Fullname,
		ID:       int64(um.ID),
	}, nil
}

func (u UserServer) Create(ctx context.Context, req *user.UserReq) (*user.Empty, error) {
	userr := &entities.User{
		Email:    req.Email,
		Fullname: req.Fullname,
	}
	err := u.userService.CreateUser(userr)
	if err != nil {
		return nil, err
	}
	return &user.Empty{}, err
}

func (u UserServer) Update(ctx context.Context, req *user.UserReq) (*user.Empty, error) {
	um := entities.User{
		Email:    req.Email,
		Fullname: req.Fullname,
	}
	err := u.userService.UpdateUser(&um)
	if err != nil {
		return nil, err
	}
	return &user.Empty{}, err
}

func (u UserServer) ListUsers(ctx context.Context, empty *user.Empty) (*user.ListUserResponse, error) {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		return nil, err
	}
	userResponses := make([]*user.UserResponse, len(users))
	for i, um := range users {
		userResponses[i] = &user.UserResponse{
			ID:       int64(um.ID),
			Email:    um.Email,
			Fullname: um.Fullname,
		}
	}
	return &user.ListUserResponse{
		Users: userResponses,
	}, nil
}

func (u UserServer) DeleteUser(ctx context.Context, req *user.UserReq) (*user.Empty, error) {
	err := u.userService.DeleteUser(int(req.GetID()))
	return &user.Empty{}, err
}
