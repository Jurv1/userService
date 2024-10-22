package main

import (
	"github.com/Jurv1/userService/database"
	grpc2 "github.com/Jurv1/userService/internal/delivery/grpc"
	repository2 "github.com/Jurv1/userService/internal/repository"
	"github.com/Jurv1/userService/internal/usecase"
	"github.com/jackc/pgx"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db := database.GetNewDB()
	defer func(db *pgx.Conn) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	ur := repository2.NewUserRepo(db)
	uq := repository2.NewUserQRepo(db)
	uu := usecase.NewUserUseCase(ur, uq)

	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpc2.NewUserServerGrpc(grpcServer, uu)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
