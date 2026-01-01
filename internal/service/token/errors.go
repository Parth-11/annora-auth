package tokenservice

import "errors"

var (
	ErrInvalidRefreshTokenFormat = errors.New("invalid refresh token format")
	ErrInvalidRefreshToken = errors.New("token has expired or has been revoked")
)
