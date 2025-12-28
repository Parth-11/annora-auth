package authservice

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Register(ctx context.Context, email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	
	return s.AuthRepo.CreateUser(ctx, email, string(hash))
}
