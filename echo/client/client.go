package main

import (
	"context"
	echo "echo/protobuf"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to Dial", err)
	}
	defer conn.Close()
	client := echo.NewEchoGRPCClient(conn)
	var msg string
	for {
		fmt.Scanln(&msg)
		if err != nil {
			fmt.Println("Failed to Read", err)
		}
		res, _ := client.Echo(context.Background(), &echo.Req{Message: msg})
		fmt.Println(res.GetMessage())
	}
}
