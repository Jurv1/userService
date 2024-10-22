package userHandler

import (
	"context"
	"fmt"
	"github.com/Jurv1/userService/internal/usecase"
	"github.com/Jurv1/userService/proto/gen/go/user/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"math/rand"
	"strconv"
)

func NewUserServerGrpc(gserver *grpc.Server, userUseCase usecase.UserUseCase) {

	userServer := &server{
		userUseCase,
		user.UnimplementedUserServiceServer{},
	}

	user.RegisterUserServiceServer(gserver, userServer)
	reflection.Register(gserver)
}

type server struct {
	usecase.UserUseCase
	user.UnimplementedUserServiceServer
}

func ConvertToJSON(message *user.CreateUserRequest) (string, error) {
	jsonBytes, err := protojson.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (s *server) GetUser(ctx context.Context, in *user.GetUserRequest) (*user.GetUserResponse, error) {
	log.Println("GetUser")
	if in.Id == 228 {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Name")
	}
	return &user.GetUserResponse{
		Id:    0,
		Name:  "Testing",
		Email: "Test",
	}, nil
}

func (s *server) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	log.Println(in)
	jsonData, err := protojson.Marshal(&user.CreateUserResponse{
		Id: strconv.Itoa(int(rand.Int63())),
	})
	if err != nil {
		log.Fatalf("failed to marshal to JSON: %v", err)
	}
	fmt.Println(string(jsonData))
	return &user.CreateUserResponse{
		Id: strconv.Itoa(228),
	}, nil
}

func (s *server) MutateUser(ctx context.Context, in *user.MutateUserRequest) (*user.MutateUserResponse, error) {
	log.Println(in)
	return &user.MutateUserResponse{
		Message: fmt.Sprintf("Test %s", in.GetInfo()),
	}, nil
}
