package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sdk-helper/logger"
	"university/cmd/proto"
	"university/handlers"
	"university/repository"
	"university/service"

	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	ServerPort    int    `envconfig:"SERVER_PORT"`
	MongoUrl      string `envconfig:"MONGO_URL"`
	MongoDatabase string `envconfig:"MONGO_DB_NAME"`
	S3_BucketName string `envconfig:"S3_BUCKET_NAME"`
	S3_Region     string `envconfig:"AWS_S3_REGION"`
	S3_AccessKey  string `envconfig:"AWS_ACCESS_KEY_ID"`
	S3_SecretKey  string `envconfig:"AWS_SECRET_ACCESS_KEY"`
	S3_Endpoint   string `envconfig:"S3_ENDPOINT_URL"`
	S3_PathStyle  string `envconfig:"S3_USE_PATH_STYLE"`
}

func main() {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err)
	}
	// Initialize MongoDB connection
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		log.Fatal(err)
	}
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ServerPort))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	defer client.Disconnect(context.Background())
	_logger := logger.NewLogger()
	db := client.Database(config.MongoDatabase)

	repository := repository.NewRepository(db)
	userService := service.NewUniversityService(repository, _logger)
	proto.RegisterUniversityServiceServer(grpcServer, &handlers.UniversityHandler{
		UnimplementedUniversityServiceServer: proto.UnimplementedUniversityServiceServer{},
		Service:                              userService,
	})
	_logger.Info("Starting the server on port :", config.ServerPort)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
