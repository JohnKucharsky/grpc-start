package main

import (
	"bufio"
	"fmt"
	userpb "github.com/JohnKucharsky/grpc-start/gen/go/user/v1"
	"google.golang.org/protobuf/proto"
	"os"
)

func main() {
	file, err := os.Open("user.bin")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	var u userpb.User
	if err := proto.Unmarshal(bytes, &u); err != nil {
		fmt.Println("error:", err.Error())
	}
	fmt.Println(u)
}
