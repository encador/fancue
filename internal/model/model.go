package model

import "github.com/golang-jwt/jwt/v5"

type ContextKey string

type CaptchaSignals struct {
	ID         string `json:"id"`
	Secret     string `json:"secret"`
	Selections []int  `json:"selection"`
}

type Signals struct {
	Captcha  CaptchaSignals
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserData struct {
	ID           string
	Username     string
	PasswordHash string
	Level        int
}

type AuthClaims struct {
	jwt.RegisteredClaims
	Username string
	UserID   string
	Level    int
}
