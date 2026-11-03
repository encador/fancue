package service

import "github.com/encador/fancue/internal/model"

type Auth struct {
	ctxKey  model.ContextKey
	signKey []byte
}
