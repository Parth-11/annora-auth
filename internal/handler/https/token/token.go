package tokenhandler

import tokenservice "github.com/AdityaTaggar05/annora-auth/internal/service/token"

type Handler struct {
	Service *tokenservice.Service
}

func NewHandler(s *tokenservice.Service) *Handler {
	return &Handler{
		Service: s,
	}
}