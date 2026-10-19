package service

type Service struct {
	Captcha CaptchaService
}

func New() *Service {
	return &Service{Captcha: CaptchaService{key: RandBytes(32), log: make(map[[32]byte]struct{}), maxLogSize: 1000000}}
}
