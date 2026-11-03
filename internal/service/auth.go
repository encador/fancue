package service

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/encador/fancue/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	ctxKey  model.ContextKey
	signKey []byte
}

func (a *Auth) GetClaims(r *http.Request) *model.AuthClaims {
	// Claims already in context
	claims, ok := r.Context().Value(a.ctxKey).(*model.AuthClaims)
	if ok {
		return claims
	}

	// Check for auth cookie
	cookie, err := r.Cookie("auth")
	if err != nil {
		return nil
	}

	// ParseToken
	claims, err = a.parseToken(cookie.Value)
	if err != nil {
		return nil
	}

	return claims
}

func (a *Auth) genToken(user model.UserData) (string, error) {

	claims := &model.AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
		Username: user.Username,
		UserID:   user.ID,
		Level:    user.Level,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(a.signKey)
}

func (a *Auth) parseToken(tString string) (*model.AuthClaims, error) {
	claims := model.AuthClaims{}
	token, err := jwt.ParseWithClaims(tString, &claims, func(t *jwt.Token) (any, error) { return a.signKey, nil })
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.AuthClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("[parseToken] invalid claims")
}

func (a *Auth) UpdateContext(ctx context.Context, claims *model.AuthClaims) context.Context {
	return context.WithValue(ctx, a.ctxKey, claims)
}

func (a *Auth) setCookie(token string, w http.ResponseWriter) error {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	return nil
}

func (a *Auth) removeCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		Expires:  time.Now(),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}
