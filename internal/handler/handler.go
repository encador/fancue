package handler

import (
	"net/http"

	"github.com/encador/fancue/internal/component"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HomePage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		component.Base().Render(r.Context(), w)
	})
}
