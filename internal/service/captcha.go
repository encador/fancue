package service

import (
	"crypto/sha256"
	"encoding/hex"
	"slices"
	"strconv"

	"github.com/a-h/templ"
	"github.com/encador/fancue/internal/component"
	"github.com/encador/fancue/internal/model"
)

var emoji = []string{
	"🍆", "🥑", "🌽", "🍒", "🥜", "🍎", "🥦",
}

// genCaptcha is used to instantiate captcha component data
//
// return choices []string, targets []int, and error
func genCaptcha(count int) ([]string, []int, error) {
	choices := make([]string, count)
	targets := []int{}
	for i := range count {
		choices[i] = emoji[RandInt(len(emoji))]
	}

	t := choices[RandInt(len(choices))]
	for i, c := range choices {
		if c == t {
			targets = append(targets, i)
		}
	}

	return choices, targets, nil
}

func getHash(arr []int, salt []byte) string {
	h := sha256.New()
	h.Write(salt[:])
	for i, n := range arr {
		if i > 0 {
			h.Write([]byte("-"))
		}
		h.Write([]byte(strconv.Itoa(n)))
	}
	return hex.EncodeToString(h.Sum(nil))
}

// Returns a fully functional captcha templ.Component
func (s *Service) NewCaptcha() templ.Component {
	icons, targets, _ := genCaptcha(2)
	return component.Captcha(icons, icons[targets[0]], getHash(targets, s.CaptchaSalt))
}

// Returns true if captcha answer is correct
func (s *Service) CaptchaCheck(signals model.CaptchaSignals) bool {
	slices.Sort(signals.Selections)
	return signals.Secret == getHash(signals.Selections, s.CaptchaSalt)
}
