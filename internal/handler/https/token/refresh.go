package tokenhandler

import (
	"encoding/json"
	"net/http"
)

type refreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) HandleRefresh(w http.ResponseWriter, r *http.Request) {
	var req refreshRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	tokens, err := h.Service.Refresh(r.Context(), req.RefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)
}
