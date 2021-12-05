package main

import (
	"log"
	"net"
	
	"grpc-category/server/category"
	"google.golang.org/grpc"
	cpb "grpc-category/proto/category"
)

func main() {
	lis, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatalf("failed to listen:%s", err)
	}

	grpcServer := grpc.NewServer()

	s:=category.Server{}

	cpb.RegisterCategoryServiceServer(grpcServer,&s)


	if err :=grpcServer.Serve(lis);err!=nil {
		log.Fatalf("failed to Serve:%s", err)
	}
}
