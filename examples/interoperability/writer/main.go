package main

import (
	userpb "github.com/JohnKucharsky/grpc-start/gen/go/user/v1"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func main() {
	u := userpb.User{
		Uuid:      "2-32-2",
		Fullname:  "John",
		BirthYear: 2800,
	}

	b, err := proto.Marshal(&u)
	if err != nil {
		log.Fatal("Marshaling error: ", err.Error())
	}

	if err := os.WriteFile("user.bin", b, 0644); err != nil {
		log.Fatal("WriteFile error: ", err.Error())
	}
}
