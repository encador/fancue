package handler

import (
	"fmt"
	"net/http"

	"github.com/encador/fancue/internal/component"
	"github.com/encador/fancue/internal/model"
	"github.com/encador/fancue/internal/service"
	"github.com/starfederation/datastar-go/datastar"
)

type Handler struct {
	s *service.Service
}

func NewHandler() *Handler {
	return &Handler{s: service.New()}
}

func (h *Handler) HomePage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		component.Base(component.Buttons()).Render(r.Context(), w)
	})
}

func (h *Handler) LoginPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// component.Base(component.LoginForm()).Render(r.Context(), w)
		component.Base(component.LoginPage(false)).Render(r.Context(), w)
	})
}

func (h *Handler) RegisterPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		component.Base(component.RegisterForm()).Render(r.Context(), w)
	})
}

func (h *Handler) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		s := model.Signals{}
		datastar.ReadSignals(r, &s)
		fmt.Println(h.s.Captcha.Validate(s.Captcha))
		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(h.s.NewCaptcha(), datastar.WithModeReplace())
	})
}
