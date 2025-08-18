package service

import (
	"context"
	"errors"
	"grpc-portal/cmd/proto"
	"http_server/models"
	"http_server/repository"
	"sdk-helper/logger"
	university "university/cmd/proto"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo             repository.Repository
	grpcClient       proto.UsersServiceClient
	universityClient university.UniversityServiceClient
	logger           *logger.Logger
}

type UserGrpcDI struct {
	Repo             repository.Repository
	Client           proto.UsersServiceClient
	UniversityClient university.UniversityServiceClient
	Logger           *logger.Logger
}

func NewService(di UserGrpcDI) UserService {
	return userService{
		repo:             di.Repo,
		logger:           di.Logger,
		grpcClient:       di.Client,
		universityClient: di.UniversityClient,
	}
}
func (u userService) CreateUser(c context.Context, user models.User) (*models.User, error) {
	if user.Password == "" || user.Name == "" || user.Email == "" {
		u.logger.ErrorContext(c, "user name or email or password is empty", user.Email, user.Password, user.Name)
		return nil, errors.New("user name or email or password is empty")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.ErrorContext(c, "error hashing password", err)
		return nil, errors.New("error hashing password")
	}
	data, err := u.grpcClient.RegisterUser(c, &proto.UserRequest{
		Name:     user.Name,
		Country:  user.Country,
		Email:    user.Email,
		Password: string(hashPassword),
		State:    user.State,
		Role:     user.Role,
		Contact:  user.Contact,
		Address:  user.Address,
	})
	if err != nil {
		u.logger.ErrorContext(c, "error creating user", err)
		return nil, err
	}
	if data.GetUuid() == "" {
		u.logger.ErrorContext(c, "uuid is empty")
		return nil, errors.New("uuid is empty")
	}
	return &models.User{
		ID:    data.Uuid,
		Name:  data.Name,
		Email: data.Email,
	}, nil
}
