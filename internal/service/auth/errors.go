package authservice

import "errors"

var (
	ErrInvalidEmailFormat = errors.New("invalid email format")
	ErrInvalidPasswordFormat = errors.New("invalid password format")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrEmailNotVerified = errors.New("email not verified")
	ErrUserNotFound = errors.New("user not found")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrTokenExpired = errors.New("token has expired")
	ErrTokenInvalid = errors.New("invalid token")
)