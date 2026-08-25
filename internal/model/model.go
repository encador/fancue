package model

type Captcha struct {
	ID         string `json:"id"`
	Secret     string `json:"secret"`
	Selections []int  `json:"selection"`
}

type Signals struct {
	Captcha  Captcha
	Username string `json:"username"`
	Password string `json:"password"`
}
