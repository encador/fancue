package handler

import (
	"fmt"
	"net/http"

	"github.com/encador/fancue/internal/captcha"
	"github.com/encador/fancue/internal/component"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HomePage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		component.Base(component.Buttons()).Render(r.Context(), w)
	})
}

func (h *Handler) TestPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		component.Base(captcha.New()).Render(r.Context(), w)
	})
}

func (h *Handler) Captcha() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
	})
}
