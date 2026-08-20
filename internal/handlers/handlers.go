package handlers

import (
	"net/http"

	"github.com/encador/fancue/internal/components"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HomePage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components.Base().Render(r.Context(), w)
	})
}
