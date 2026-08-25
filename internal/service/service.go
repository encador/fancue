package service

type Service struct {
	CaptchaSalt []byte
}

func New() *Service {
	return &Service{CaptchaSalt: RandBytes(32)}
}
