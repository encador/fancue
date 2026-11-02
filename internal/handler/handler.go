package handler

import (
	"net/http"
	"strconv"
	"time"

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

func DisplayMsg(sse *datastar.ServerSentEventGenerator, message, color string) {
	switch color {
	case "blue", "green", "orange", "red":
	default:
		color = "blue"
	}

	sse.PatchElementTempl(component.MsgBox(message, color), datastar.WithSelectorID("msg-bar"), datastar.WithModeAppend())
}

func (h *Handler) HomePage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		component.Base(component.Buttons()).Render(r.Context(), w)
	})
}

func (h *Handler) Msg() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)
		DisplayMsg(sse, "Seconds: "+strconv.Itoa(time.Now().Second()), "blue")
	})
}

func (h *Handler) LoginPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// component.Base(component.LoginForm()).Render(r.Context(), w)
		component.Base(component.LoginPage(false), h.s.Captcha.New()).Render(r.Context(), w)
	})
}

func (h *Handler) RegisterPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// component.Base(component.RegisterForm()).Render(r.Context(), w)
		component.Base(component.LoginPage(true), h.s.Captcha.New()).Render(r.Context(), w)
	})
}

func (h *Handler) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		s := model.Signals{}
		datastar.ReadSignals(r, &s)
		sse := datastar.NewSSE(w, r)
		captchaValid := h.s.Captcha.Validate(s.Captcha)
		if !captchaValid {
			DisplayMsg(sse, "Invalid Captcha", "red")
			sse.PatchElementTempl(h.s.Captcha.New(), datastar.WithSelectorID("captcha"), datastar.WithModeReplace())
			return
		}
		DisplayMsg(sse, "Captcha Success", "green")
	})
}
