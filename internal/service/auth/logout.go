package authservice

import (
	"context"

	tokenservice "github.com/AdityaTaggar05/annora-auth/internal/service/token"
)

func (s *Service) Logout(ctx context.Context, oldToken string) error {
	if !tokenservice.IsValidRefreshToken(oldToken) {
		return tokenservice.ErrInvalidRefreshTokenFormat
	}

	return s.TokenRepo.RevokeRefreshToken(ctx, oldToken)
}
