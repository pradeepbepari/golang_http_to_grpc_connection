package service

import (
	"context"
	"fmt"
	"http_server/repository"
	"io"
	"sdk-helper/aws"
	"sdk-helper/logger"
)

type fileService struct {
	repo     repository.Repository
	awsStore aws.AwsStoreService
	logger   *logger.Logger
}

func NewFileService(repo repository.Repository, awsStore aws.AwsStoreService, logger *logger.Logger) FileService {
	return &fileService{
		repo:     repo,
		awsStore: awsStore,
		logger:   logger,
	}
}
func (s *fileService) FileUploadToS3(ctx context.Context, reader io.Reader, key string) (string, error) {
	url, err := s.awsStore.Upload(ctx, reader, key)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to upload file to s3", err)
		return "", fmt.Errorf("error uploading file to s3 %v", err.Error())
	}
	return url, nil
}
