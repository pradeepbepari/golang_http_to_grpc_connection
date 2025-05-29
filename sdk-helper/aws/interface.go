package aws

import (
	"context"
	"io"
)

type AwsStoreService interface {
	Upload(context.Context, io.Reader, string) (string, error)
	GeneratePresignedUploadURL(ctx context.Context, key string) (string, error)
}
