package grpc

import (
	"bm-users/src/apps/grpc/proto/user"
	"bm-users/src/apps/grpc/servers"
	"bm-users/src/domain/repositories/user/userGorm"
	"bm-users/src/domain/services"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net"
)

func RunServer() {
	l, err := net.Listen("tcp", "localhost:8101")
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	dsn := "host=localhost user=postgres password=postgres dbname=bm_users port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	userRepo := userGorm.NewGormUserRepository(db)
	userService := services.NewUserService(userRepo, logger)
	userServer := servers.NewUserServer(userService)
	user.RegisterUserServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	err = grpcServer.Serve(l)
	if err != nil {
		panic(err)
	}
}
