package main

import (
	"context"
	"fmt"
	"grpc-portal/handler"
	"log"
	"net"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"grpc-portal/cmd/proto"
	"grpc-portal/repository"
	"grpc-portal/service"
	"sdk-helper/logger"
)

func main() {
	mongoURI := os.Getenv("MONGO_URL")
	mongoDBName := os.Getenv("MONGO_DB_NAME")
	serverPort := os.Getenv("GRPC_SERVER_PORT")
	missingVars := []string{}

	if mongoURI == "" {
		missingVars = append(missingVars, "MONGO_URL")
	}
	if serverPort == "" {
		missingVars = append(missingVars, "GRPC_SERVER_PORT")
	}
	if serverPort == "" {
		missingVars = append(missingVars, "MONGO_DB_NAME")
	}
	if len(missingVars) > 0 {
		log.Fatalf("The following environment variables are missing: %v", missingVars)
	}

	ctx := context.Background()
	_logger := logger.NewLogger()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	db := client.Database(mongoDBName)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}
	_logger.Info("Starting the server on port %s", serverPort)
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	repo := repository.NewRepository(db)
	userService := service.NewUserService(repo, _logger)
	proto.RegisterUsersServiceServer(grpcServer, &handler.Server{
		UnimplementedUsersServiceServer: proto.UnimplementedUsersServiceServer{},
		UserService:                     userService,
		Logger:                          _logger,
	})
	_logger.Info("server is running")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
