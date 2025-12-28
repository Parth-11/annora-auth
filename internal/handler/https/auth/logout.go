package authhandler

import (
	"encoding/json"
	"net/http"
)

type logoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	var req logoutRequest
	
	json.NewDecoder(r.Body).Decode(&req)

    err := h.Service.Logout(r.Context(), req.RefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
    w.WriteHeader(http.StatusNoContent)
}
