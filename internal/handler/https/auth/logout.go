package authhandler

import (
	"encoding/json"
	"net/http"

	tokenservice "github.com/AdityaTaggar05/annora-auth/internal/service/token"
)

type logoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	var req logoutRequest
	
	json.NewDecoder(r.Body).Decode(&req)

    if err := h.Service.Logout(r.Context(), req.RefreshToken); err != nil {
		switch err {
			case tokenservice.ErrInvalidRefreshTokenFormat:
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}
    w.WriteHeader(http.StatusNoContent)
}
