package tokenservice

func IsValidRefreshToken(token string) bool {
	return len(token) == 32
}
