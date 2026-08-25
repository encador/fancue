package service

import "fmt"

type Service struct {
	CaptchaSalt []byte
}

func New() *Service {
	salt, err := getSalt(32)
	if err != nil {
		fmt.Println("[ERROR]: captcha salt gen failed")
	}
	return &Service{CaptchaSalt: salt}
}
