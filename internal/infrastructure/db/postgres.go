package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AdityaTaggar05/annora-auth/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(ctx context.Context, cfg config.PostgresConfig) *pgxpool.Pool {
	pg_cfg, err := pgxpool.ParseConfig(cfg.URL)
    if err != nil {
        log.Fatal("[ERR] Unable to connect to database: ", err)
    }

    pg_cfg.MaxConns = int32(cfg.MaxOpenConns)
    pg_cfg.MinConns = 5
    pg_cfg.MaxConnLifetime = time.Hour

    DB, err := pgxpool.NewWithConfig(ctx, pg_cfg)

    if err != nil {
		log.Fatal("[ERR] Unable to connect to database: ", err)
	}

	err = DB.Ping(ctx)
	if err != nil {
		log.Fatal("[ERR] Could not ping the database: ", err)
	}
	fmt.Println("[DEBUG] Connected to the database!")

    return DB
}
