package main

import (
	"bm-users/src/apps/gin"
	"bm-users/src/apps/grpc"
)

func main() {
	go gin.RunServer()
	grpc.RunServer()
}
