package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"slices"
	"strconv"

	"github.com/a-h/templ"
	"github.com/encador/fancue/internal/component"
	"github.com/encador/fancue/internal/model"
)

var emoji = []string{
	"🍆", "🥑", "🌽", "🍒", "🥜", "🍎", "🥦",
}

func getSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	return salt, err
}

// genCaptcha is used to instantiate captcha component data
//
// return choices []string, targets []int, and error
func genCaptcha(count int) ([]string, []int, error) {
	choices := make([]string, count)
	targets := []int{}
	for i := range count {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(emoji))))
		choices[i] = emoji[int(num.Int64())]
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(choices))))
	if err != nil {
		return []string{}, []int{}, err
	}
	t := choices[n.Int64()]
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
	fmt.Println(getHash(targets, s.CaptchaSalt))
	return component.Captcha(icons, icons[targets[0]], getHash(targets, s.CaptchaSalt))
}

// Returns true if captcha answer is correct
func (c *Service) CaptchaCheck(s model.CaptchaSignals) bool {
	slices.Sort(s.Selections)
	return s.Secret == getHash(s.Selections, c.CaptchaSalt)
}
