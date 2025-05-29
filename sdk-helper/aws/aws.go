package aws

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type awsConfig struct {
	s3Client        *s3.Client
	preSignedClient *s3.PresignClient
	bucketName      string
}

func NewAwsConfig(c aws.Config, bucket string) AwsStoreService {
	s3Client := s3.NewFromConfig(c, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	preSignedClient := s3.NewPresignClient(s3Client)
	return &awsConfig{
		s3Client:        s3Client,
		preSignedClient: preSignedClient,
		bucketName:      *aws.String(bucket),
	}
}

func (s *awsConfig) Upload(ctx context.Context, body io.Reader, key string) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buf.Bytes())

	_, err := s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &s.bucketName,
		Key:         &key,
		Body:        bytes.NewReader(buf.Bytes()),
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.bucketName, key)
	return url, nil
}
func (s *awsConfig) GeneratePresignedUploadURL(ctx context.Context, key string) (string, error) {
	req, err := s.preSignedClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.bucketName,
		Key:    &key,
		ACL:    types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return "", err
	}
	return req.URL, nil
}
