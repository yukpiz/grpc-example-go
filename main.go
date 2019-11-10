package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/yukpiz/grpc-example-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	server := grpc.NewServer()

	svc := &Service{}
	pb.RegisterExampleServiceServer(server, svc)
	reflection.Register(server)

	conn, err := net.Listen("tcp", ":1111")
	if err != nil {
		fmt.Printf("network I/O error: %v", err)
		os.Exit(1)
	}

	fmt.Println("...Waiting for localhost:1111")
	if err := server.Serve(conn); err != nil {
		fmt.Printf("serve error: %v", err)
		os.Exit(1)
	}
}

type Service struct{}

func (*Service) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	fmt.Println("Helloが呼び出されました！")

	return &pb.HelloResponse{
		Id:   req.Id,
		Name: req.Name,
	}, nil
}
