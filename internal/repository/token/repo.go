package tokenrepo

import "github.com/jackc/pgx/v5/pgxpool"

type TokenRepository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *TokenRepository {
	return &TokenRepository{
		DB: db,
	}
}
