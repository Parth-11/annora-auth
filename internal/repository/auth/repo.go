package authrepo

import "github.com/jackc/pgx/v5/pgxpool"

type AuthRepository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}
