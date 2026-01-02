package authhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	authservice "github.com/AdityaTaggar05/annora-auth/internal/service/auth"
)

func (h *Handler) HandleResendVerification(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.ResendVerification(r.Context(), req.Email); err != nil {
		switch err {
			case authservice.ErrInvalidEmailFormat:
				http.Error(w, err.Error(), http.StatusBadRequest)
			case authservice.ErrUserNotFound:
				http.Error(w, err.Error(), http.StatusNotFound)
			case authservice.ErrTooManyRequests:
				http.Error(w, err.Error(), http.StatusTooManyRequests)
			default:
				fmt.Println(err.Error())
				http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}