package server

import (
	"github.com/Jurv1/userService/proto/gen/go/user/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	user.userv1
}

func NewServer() *Server {
	return &Server{}
}

func (s Server) Run() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	user.userv1.RegisterUserServiceServer(grpcServer, NewServer())

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
