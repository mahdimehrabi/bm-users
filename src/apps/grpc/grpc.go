package grpc

import (
	"bm-users/src/apps/grpc/proto/user"
	"bm-users/src/apps/grpc/servers"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func RunServer() {
	l, err := net.Listen("tcp", "localhost:8101")
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	user.RegisterUserServer(grpcServer, &servers.UserServer{})
	reflection.Register(grpcServer)

	fmt.Println("runnnn")
	err = grpcServer.Serve(l)
	if err != nil {
		panic(err)
	}
}
