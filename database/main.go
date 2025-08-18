package main

import (
	"_database/migrations/migrate"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type Config struct {
	DatabaseHost     string `envconfig:"DB_HOST"`
	DatabasePort     string `envconfig:"DB_PORT"`
	DatabaseUser     string `envconfig:"DB_USER"`
	DatabasePassword string `envconfig:"DB_PASSWORD"`
	DatabaseName     string `envconfig:"DB_NAME"`
	AppPort          string `envconfig:"APP_PORT"`
	MigrationsPath   string `envconfig:"MIGRATIONS_PATH" default:"migrations"`
}

func main() {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("error processing env variables: %v", err)
	}

	log.Println("Configuration loaded successfully:", config)

	// Database connect
	connection, err := ConnectToDatabase(
		config.DatabaseHost, config.DatabasePort,
		config.DatabaseUser, config.DatabasePassword, config.DatabaseName,
	)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer connection.Close()

	// Run migration UP on start
	if err := migrate.RunMigration(migrate.Config{
		DBHost:     config.DatabaseHost,
		DBPort:     config.DatabasePort,
		DBUser:     config.DatabaseUser,
		DBPassword: config.DatabasePassword,
		DBName:     config.DatabaseName,
	}, "up"); err != nil {
		log.Fatalf("migration up failed: %v", err)
	}

	// Handle shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Println("Shutting down... Running migration down")
		if err := migrate.RunMigration(migrate.Config{
			DBHost:     config.DatabaseHost,
			DBPort:     config.DatabasePort,
			DBUser:     config.DatabaseUser,
			DBPassword: config.DatabasePassword,
			DBName:     config.DatabaseName,
		}, "down"); err != nil {
			log.Printf("migration down failed: %v", err)
		}
		os.Exit(0)
	}()

	// Start server
	router := NewRouter()
	if err := router.Run(fmt.Sprintf(":%s", config.AppPort)); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}

func ConnectToDatabase(host, port, user, password, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	log.Println("Successfully connected to the database")
	return db, nil
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})
	return router
}
