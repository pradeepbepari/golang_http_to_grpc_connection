package handler

import (
	"grpc-portal/cmd/proto"
	"grpc-portal/service"
	"sdk-helper/logger"
	university "university/cmd/proto"
)

type Server struct {
	proto.UnimplementedUsersServiceServer
	UserService       service.UserService
	UniversityService university.UniversityServiceClient
	Logger            *logger.Logger
}
