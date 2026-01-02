package authservice

import (
	"context"
)

func (s *Service) ResendVerification(ctx context.Context, email string) error {
	if !isValidEmail(email) {
		return ErrInvalidEmailFormat
	}

	user, err := s.AuthRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return ErrUserNotFound
	}

	key := "email_resend:" + email
	count, _ := s.TokenRepo.RDB.Incr(ctx, key).Result()

	if count == 1 {
		s.TokenRepo.RDB.Expire(ctx, key, s.Mailer.ResendLimitTTL)
	}

	if count > int64(s.Mailer.ResendLimit) {
		return ErrTooManyRequests
	}

	token, err := generateEmailVerificationToken()
	if err != nil {
		return err
	}

	key = "email_verify:" + token

	err = s.TokenRepo.CreateEmailToken(ctx, key, user.ID, s.EmailTokenTTL)
	if err != nil {
		return err
	}

	return s.Mailer.SendVerificationEmail(email, token)
}