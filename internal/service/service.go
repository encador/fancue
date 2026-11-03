package service

type Service struct {
	Captcha *Captcha
	User    *User
	Auth    *Auth
}

func New() *Service {

	service := &Service{}
	captcha := &Captcha{key: RandBytes(32), log: make(map[[32]byte]struct{}), maxLogSize: 100000}
	service.Captcha = captcha

	user := &User{}
	service.User = user

	auth := &Auth{ctxKey: "auth", signKey: RandBytes(32)}
	service.Auth = auth

	return service
}
