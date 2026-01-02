package authservice

import (
	"context"
	"fmt"
)

func (s *Service) ResendVerification(ctx context.Context, email string) error {
	if !isValidEmail(email) {
		return ErrInvalidEmailFormat
	}

	user, err := s.AuthRepo.GetUserByEmail(ctx, email)
	fmt.Println("FIRST")
	if err != nil {
		return ErrUserNotFound
	}

	key := "email_resend:" + email
	count, _ := s.TokenRepo.RDB.Incr(ctx, key).Result()

	if count == 1 {
		s.TokenRepo.RDB.Expire(ctx, key, s.Mailer.ResendLimitTTL)
	}

	fmt.Println("SECOND")
	if count > int64(s.Mailer.ResendLimit) {
		return ErrTooManyRequests
	}

	token, err := generateEmailVerificationToken()
	fmt.Println("THIRD")
	if err != nil {
		return err
	}

	fmt.Println(user.ID)
	key = "email_verify:" + token

	err = s.TokenRepo.CreateEmailToken(ctx, key, user.ID, s.EmailTokenTTL)
	fmt.Println("FOURTH")
	if err != nil {
		return err
	}

	fmt.Println("FIFTH")
	return s.Mailer.SendVerificationEmail(email, token)
}