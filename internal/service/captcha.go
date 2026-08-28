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

type Captcha struct {
	salt []byte
	log  map[[32]byte]struct{}
}

// genCaptcha is used to instantiate captcha component data
//
// return choices []string, targets []int, and error
func genCaptcha(count int) ([]string, []int) {
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

	return choices, targets
}

func hashCaptcha(id string, selections []int, salt []byte) string {
	h := sha256.New()
	h.Write(salt[:])
	for i, n := range selections {
		if i > 0 {
			h.Write([]byte("-"))
		}
		h.Write([]byte(strconv.Itoa(n)))
	}
	h.Write([]byte(id))
	return hex.EncodeToString(h.Sum(nil))
}

// Returns a fully functional captcha templ.Component
func (s *Service) NewCaptcha() templ.Component {
	icons, targets := genCaptcha(10)
	id := RandString()
	hash := hashCaptcha(id, targets, s.Captcha.salt)

	return component.Captcha(id, icons, icons[targets[0]], hash)
}

// Check if captcha hash was previously Seen
//
// Set seen to true if not
func (c *Captcha) Seen(hash string) bool {
	tmp, err := hex.DecodeString(hash)
	if err != nil {
		// technically not seen
		return true
	}
	b := [32]byte(tmp)
	_, ok := c.log[b]
	if ok {
		return true
	}
	c.log[b] = struct{}{}
	return false
}

// Returns true if captcha answer is correct
//
// Returns false if captcha was previously validated
func (c *Captcha) Validate(signals model.Captcha) bool {
	// Check if captcha already submitted before
	if signals.Secret == "" || c.Seen(signals.Secret) {
		return false
	}

	slices.Sort(signals.Selections)
	if signals.Secret != hashCaptcha(signals.ID, signals.Selections, c.salt) {
		return false
	}

	return true
}
