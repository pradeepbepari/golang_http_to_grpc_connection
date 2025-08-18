package migrate

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	MigrationPath string
}

func RunMigration(cfg Config, direction string) error {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	migrationsPath := "file:///app/migrations"

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("migration up failed: %w", err)
		}
		log.Println("Migration up applied successfully")

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("migration down failed: %w", err)
		}
		log.Println("Migration down applied successfully")
	}

	return nil
}
