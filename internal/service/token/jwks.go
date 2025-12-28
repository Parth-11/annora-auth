package tokenservice

func (s *Service) JWKS() map[string]any {
	jwk := s.SigningKey.PublicKeyToJWK()

	return map[string]any{
		"keys": []any{jwk},
	}
}
