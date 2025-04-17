package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

func NewPostgresConnection() *pgxpool.Pool {
	dsn := os.Getenv("DATABASE_URL")
	conn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return conn

}
