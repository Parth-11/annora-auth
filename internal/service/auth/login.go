package authservice

import (
	"context"
	"time"

	"github.com/AdityaTaggar05/annora-auth/internal/model"
	"github.com/AdityaTaggar05/annora-auth/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, email, password string) (model.TokenPair, error) {
	tokens := model.TokenPair{}

	user, err := s.AuthRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return tokens, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return tokens, err
	}

	tokens.AccessToken, err = utils.GenerateJWT(user.ID, s.SigningKey, s.Config.AccessTTL)
	if err != nil {
		return tokens, err
	}

	tokens.RefreshToken, err = utils.GenerateRefreshToken()
	if err != nil {
		return tokens, err
	}

	err = s.TokenRepo.CreateRefreshToken(ctx, user.ID, tokens.RefreshToken, time.Now().Add(s.Config.RefreshTTL))
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}
