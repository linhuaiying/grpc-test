package main

import (
	"context"
	"fmt"
	hello_grpc "go.mod/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:8888", grpc.WithInsecure())
	defer conn.Close()
	client := hello_grpc.NewHelloWordServerClient(conn)
	rsp, _ := client.GetHelloWord(context.Background(), &hello_grpc.HelloWordRequest{})
	fmt.Println(rsp)
}
