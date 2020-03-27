package protoserver

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func Serve() {
	// リッスン処理
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// サーバ起動
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
