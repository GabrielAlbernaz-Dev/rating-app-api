package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	dsn := getDatabaseURL()

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v\n", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.HealthCheckPeriod = 5 * time.Minute

	DB, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Database connection established")
}

func Close() {
	if DB != nil {
		DB.Close()
		fmt.Println("Database connection closed")
	}
}

func getDatabaseURL() string {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Environment variable DATABASE_URL is not set")
	}
	return dbURL
}
