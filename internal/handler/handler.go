package handler

import (
	"fmt"
	"net/http"

	"github.com/encador/fancue/internal/captcha"
	"github.com/encador/fancue/internal/component"
	"github.com/starfederation/datastar-go/datastar"
)

type Signals struct {
	Captcha  captcha.Signals
	Username string `json:"username"`
	Password string `json:"password"`
}

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
		challenge := captcha.New()
		component.Base(component.TestForm(challenge)).Render(r.Context(), w)
	})
}

func (h *Handler) LoginPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		component.Base(component.LoginForm()).Render(r.Context(), w)
	})
}

func (h *Handler) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		s := Signals{}
		datastar.ReadSignals(r, &s)
		fmt.Println(s.Username, s.Password)
	})
}
