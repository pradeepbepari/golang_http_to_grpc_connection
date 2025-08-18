package repository

import (
	"database/sql"
	"sdk-helper/logger"
)

type repo struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewRepository(db *sql.DB, logger *logger.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}
