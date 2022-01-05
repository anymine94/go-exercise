package main

import (
	"context"
	echo "echo/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	echo.UnimplementedEchoGRPCServer
}

func (S *server) Echo(ctx context.Context, req *echo.Req) (res *echo.Res, err error) {
	msg := req.GetMessage()
	return &echo.Res{Message: msg}, nil
}
func main() {
	listen, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		log.Fatalf("failed to listen", err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoGRPCServer(s, &server{})
	s.Serve(listen)
}
