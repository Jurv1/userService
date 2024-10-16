package main

import (
	"context"
	"fmt"
	userv1 "github.com/Jurv1/userService/proto/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"math/rand"
	"net"
	"strconv"
)

type server struct {
	userv1.UnimplementedUserServiceServer
}

func ConvertToJSON(message *userv1.CreateUserRequest) (string, error) {
	jsonBytes, err := protojson.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (s *server) GetUser(ctx context.Context, in *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	log.Println("GetUser")
	if in.Id == 228 {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Name")
	}
	return &userv1.GetUserResponse{
		Id:    0,
		Name:  "Testing",
		Email: "Test",
	}, nil
}

func (s *server) CreateUser(ctx context.Context, in *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	log.Println(in)
	jsonData, err := protojson.Marshal(&userv1.CreateUserResponse{
		Id: strconv.Itoa(int(rand.Int63())),
	})
	if err != nil {
		log.Fatalf("failed to marshal to JSON: %v", err)
	}
	fmt.Println(string(jsonData))
	return &userv1.CreateUserResponse{
		Id: strconv.Itoa(228),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userv1.RegisterUserServiceServer(grpcServer, &server{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
