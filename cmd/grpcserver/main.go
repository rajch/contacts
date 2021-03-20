package main

import (
	"fmt"
	"log"
	"net"

	cgrpc "github.com/rajch/contacts/pkg/grpc"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8085")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Setting up contacts grpc server...")
	grpcServer := grpc.NewServer()
	cgrpc.RegisterContactServiceServer(grpcServer, &contactServer{})

	fmt.Println("Starting grpc server...")
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Printf("Server stopped with error: %v.", err)
	}
}
