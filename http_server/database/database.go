package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func ConnectionDatabase(c *mysql.Config) (*sql.DB, error) {
	connection, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := connection.Ping(); err != nil {
		return nil, err
	}
	return connection, nil
}
