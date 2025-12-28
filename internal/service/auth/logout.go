package authservice

import "context"

func (s *Service) Logout(ctx context.Context, oldToken string) error {
	return s.TokenRepo.RevokeRefreshToken(ctx, oldToken)
}
