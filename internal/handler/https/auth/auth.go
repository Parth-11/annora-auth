package authhandler

import (
	authservice "github.com/AdityaTaggar05/annora-auth/internal/service/auth"
)

type Handler struct {
	Service *authservice.Service
}

func NewHandler(s *authservice.Service) *Handler {
	return &Handler{
		Service: s,
	}
}
