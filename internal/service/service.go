package service

type Service struct {
	Captcha *Captcha
	User    *User
}

func New() *Service {

	service := &Service{}
	captcha := &Captcha{key: RandBytes(32), log: make(map[[32]byte]struct{}), maxLogSize: 100000}
	service.Captcha = captcha

	user := &User{}
	service.User = user

	return service
}
