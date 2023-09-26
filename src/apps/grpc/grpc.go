package grpc

import (
	"bm-users/src/apps/grpc/proto/user"
	"bm-users/src/apps/grpc/servers"
	"google.golang.org/grpc"
	"net"
)

func RunServer() {
	l, err := net.Listen("tcp", "localhost:8101")
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	user.RegisterUserServer(grpcServer, servers.UserServer{})
	err = grpcServer.Serve(l)
	if err != nil {
		panic(err)
	}
}
