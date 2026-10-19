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
	key        []byte
	log        map[[32]byte]struct{}
	maxLogSize int
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
func (s *Service) New() templ.Component {
	icons, targets := genCaptcha(10)
	id := RandString()
	hash := hashCaptcha(id, targets, s.Captcha.key)

	return component.Captcha(id, icons, icons[targets[0]], hash)
}

// Check if captcha hash was previously submitted
func (c *Captcha) IsNew(hash [32]byte) bool {
	if _, ok := c.log[hash]; ok {
		return false
	}
	return true
}

// Adds a given captcha hash to the log
//
// Rotates the secret key when the log exceeds max size.
// Currenty this invalidats all previously issued captchas
func (c *Captcha) AddToLog(hash [32]byte) {
	c.log[hash] = struct{}{}
	if len(c.log) > c.maxLogSize {
		c.key = RandBytes(32)
		c.log = make(map[[32]byte]struct{})
	}
	return
}

// Returns true if captcha answer is correct
//
// Returns false if captcha was previously validated
func (c *Captcha) Validate(signals model.CaptchaSignals) bool {
	if signals.ID == "" || signals.Secret == "" || len(signals.Selections) == 0 {
		return false
	}

	tmp, err := hex.DecodeString(signals.Secret)
	if err != nil {
		return false
	}
	hash := [32]byte(tmp)

	if signals.Secret == "" || !c.IsNew(hash) {
		c.AddToLog(hash)
		return false
	}
	c.AddToLog(hash)

	slices.Sort(signals.Selections)
	if signals.Secret != hashCaptcha(signals.ID, signals.Selections, c.key) {
		return false
	}

	return true
}
