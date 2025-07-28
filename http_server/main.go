package main

import (
	"context"
	"fmt"
	"grpc-portal/cmd/proto"
	"http_server/database"
	handlers "http_server/handler"
	"http_server/repository"
	"http_server/routes"
	"http_server/service"
	"log"
	"time"

	university "university/cmd/proto"

	"github.com/aws/aws-sdk-go-v2/aws"
	_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "http_server/docs"
	_aws "sdk-helper/aws"
	"sdk-helper/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	DatabaseHost     string `envconfig:"DB_HOST"`
	DatabasePort     string `envconfig:"DB_PORT"`
	DatabaseUser     string `envconfig:"DB_USER"`
	DatabasePassword string `envconfig:"DB_PASSWORD"`
	DatabaseName     string `envconfig:"DATABASE"`
	ServerPort       string `envconfig:"HTTP_SERVER_PORT"`
	GrpcPort         string `envconfig:"GRPC_SERVER_PORT"`
	UniversityPort   string `envconfig:"UNIVERSITY_SERVER_PORT"`
	S3_BucketName    string `envconfig:"S3_BUCKET_NAME"`
	S3_Region        string `envconfig:"AWS_S3_REGION"`
	S3_AccessKey     string `envconfig:"AWS_ACCESS_KEY_ID"`
	S3_SecretKey     string `envconfig:"AWS_SECRET_ACCESS_KEY"`
	S3_Endpoint      string `envconfig:"S3_ENDPOINT_URL"`
	S3_PathStyle     string `envconfig:"S3_USE_PATH_STYLE"`
}

func main() {

	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("error processing env variables: %v", err)
	}
	db, err := database.ConnectionDatabase(&mysql.Config{
		User:                 config.DatabaseUser,
		Passwd:               config.DatabasePassword,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", config.DatabaseHost, config.DatabasePort),
		DBName:               config.DatabaseName,
		AllowNativePasswords: true,
	})
	if err != nil {
		log.Fatalf("error connecting mysql database %v", err)
	}
	ctx := context.Background()
	_logger := logger.NewLogger()
	awsConfig, err := _config.LoadDefaultConfig(ctx, _config.WithRegion(config.S3_Region),
		_config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.S3_AccessKey, config.S3_SecretKey, "")),
		_config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           config.S3_Endpoint, // or localhost:9000 if not using Docker DNS
				SigningRegion: config.S3_Region,
			}, nil
		})),
	)
	if err != nil {
		log.Fatal(err)
	}
	awsStorage := _aws.NewAwsConfig(awsConfig, config.S3_BucketName)
	if err != nil {
		log.Fatalf("error loading aws config %v", err)
	}
	_grpcPortalConn, err := grpc.NewClient(config.GrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error connecting grpc server %v", err)
	}
	_universityConn, err := grpc.NewClient(config.UniversityPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error connecting grpc server %v", err)
	}
	database.Schemachange(ctx, db)
	repository := repository.NewRepository(db)
	authService := service.NewAuthService(repository, _logger)
	fileService := service.NewFileService(repository, awsStorage, _logger)
	userService := service.NewService(service.UserGrpcDI{
		Repo:             repository,
		Client:           proto.NewUsersServiceClient(_grpcPortalConn),
		UniversityClient: university.NewUniversityServiceClient(_universityConn),
		Logger:           _logger,
	})
	authHandler := handlers.NewAuthHandler(authService, _logger)
	userHandler := handlers.NewHandler(userService, _logger)
	awsHandler := handlers.NewFileHandler(fileService, _logger)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Geo-Location", "X-Language", "X-Timezone"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	dependencies := routes.Dependencies{
		AuthHandler:  authHandler,
		UserHandler:  userHandler,
		FileHandlers: awsHandler,
	}
	routes.ApiRoutes(dependencies, router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	defer db.Close()
	defer _grpcPortalConn.Close()
	defer _universityConn.Close()
	router.Run(fmt.Sprintf(":%s", config.ServerPort))

}
