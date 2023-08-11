package main

import (
	"context"
	hello_grpc "github.com/linhuaiying/grpc-test/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloWordServerServer
}

func (s *server) GetHelloWord(context.Context, *hello_grpc.HelloWordRequest) (*hello_grpc.HelloWordReply, error) {
	reply := hello_grpc.HelloWordReply{
		Retcode: 0,
		Retmsg:  "请求成功",
	}
	return &reply, nil
}

func main() {
	l, _ := net.Listen("tcp", "8888")
	s := grpc.NewServer()
	hello_grpc.RegisterHelloWordServerServer(s, &server{})
	s.Serve(l)
}
