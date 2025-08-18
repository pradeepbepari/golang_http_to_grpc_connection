package service

import (
	"context"
	"http_server/models"
	"io"
)

type UserService interface {
	CreateUser(context.Context, models.User) (*models.User, error)
}
type FileService interface {
	FileUploadToS3(context.Context, io.Reader, string) (string, error)
}
type AuthService interface {
	LoginUser(context.Context, string) (models.User, error)
}
