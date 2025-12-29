package authhandler

import "net/http"

func (h *Handler) HandleResendVerification(w http.ResponseWriter, r *http.Request) {
	h.Service.Mailer.SendVerificationEmail("", "sample-token")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("verification email sent"))
}