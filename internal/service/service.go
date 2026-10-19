package service

type Service struct {
	Captcha Captcha
}

func New() *Service {
	return &Service{Captcha: Captcha{key: RandBytes(32), log: make(map[[32]byte]struct{}), maxLogSize: 100000}}
}
