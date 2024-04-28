package main

import (
	"context"
	userpb "github.com/JohnKucharsky/grpc-start/gen/go/user/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type userService struct {
	userpb.UnsafeUserServiceServer
}

func (u *userService) GetUser(ctx context.Context, request *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{User: &userpb.User{
		Uuid:      request.Uuid,
		Fullname:  "sdfgsdfg",
		BirthYear: 0,
		Salary:    0,
		Addresses: nil,
	}}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &userService{})
	grpcServer.Serve(lis)
}
