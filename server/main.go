package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatalf("failed to listen:%s", err)
	}

	grpcServer := grpc.NewServer()

	if err :=grpcServer.Serve(lis);err!=nil {
		log.Fatalf("failed to Serve:%s", err)
	}
}
