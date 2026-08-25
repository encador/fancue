package model

type CaptchaSignals struct {
	Secret     string `json:"secret"`
	Selections []int  `json:"selection"`
}

type Signals struct {
	Captcha  CaptchaSignals
	Username string `json:"username"`
	Password string `json:"password"`
}
