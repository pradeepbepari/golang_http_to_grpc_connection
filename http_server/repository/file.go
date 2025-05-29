package repository

import (
	"context"
	"io"
)

func (r *repo) UploadFile(context.Context, io.Reader, string) (string, error) {
	return "", nil
}
