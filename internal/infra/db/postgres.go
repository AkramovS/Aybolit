package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

func InitPostgres() *pgxpool.Pool {
	dsn := "postgres://postgres:Akramchik938747405@localhost:5432/aybolit?sslmode=disable"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("❌ Could not connect to Postgres: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("❌ Could not ping Postgres: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL!")
	return pool
}
