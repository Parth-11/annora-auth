package tokenservice

import (
	"github.com/AdityaTaggar05/annora-auth/internal/config"
	"github.com/AdityaTaggar05/annora-auth/internal/model"
	tokenrepo "github.com/AdityaTaggar05/annora-auth/internal/repository/token"
)

type Service struct {
	TokenRepo *tokenrepo.TokenRepository
	Config config.JWTConfig
	SigningKey *model.SigningKey
}

func NewService(tokenRepo *tokenrepo.TokenRepository, cfg config.JWTConfig, signingKey *model.SigningKey) *Service {
	return &Service{
		TokenRepo: tokenRepo,
		Config: cfg,
		SigningKey: signingKey,
	}
}
