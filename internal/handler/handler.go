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
		component.Base(component.LoginForm(), h.s.NewCaptcha() ).Render(r.Context(), w)
		fmt.Println(service.RandInt(100))
	})
}

func (h *Handler) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		s := model.Signals{}
		datastar.ReadSignals(r, &s)
		fmt.Println(s.Username, s.Password)
		fmt.Println(s.Captcha.Selections)
	})
}
