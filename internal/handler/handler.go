package handler

import (
	"net/http"

	"github.com/encador/fancue/internal/captcha"
	"github.com/encador/fancue/internal/component"
	"github.com/starfederation/datastar-go/datastar"
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
		challenge := captcha.New()
		component.Base(component.TestForm(challenge)).Render(r.Context(), w)
	})
}

type PageSignals struct {
	Captcha captcha.Signals
}

func (h *Handler) Captcha() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signals := PageSignals{}
		datastar.ReadSignals(r, &signals)
		sse := datastar.NewSSE(w, r)
		if !captcha.Check(signals.Captcha) {
			sse.PatchElementTempl(captcha.New())

		}
	})
}
