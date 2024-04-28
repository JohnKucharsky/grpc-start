package main

import (
	"context"
	"fmt"
	userpb "github.com/JohnKucharsky/grpc-start/gen/go/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9879", opts...)
	if err != nil {
		log.Fatal("fail to dial:", err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{
		Uuid: "any",
	})
	if err != nil {
		log.Fatal("fail to get user:", err.Error())
	}
	fmt.Println(res)
}
