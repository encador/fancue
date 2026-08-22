package captcha

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/a-h/templ"
	"github.com/encador/fancue/internal/component"
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

// Returns a fully functional captcha templ.Component
func New() templ.Component {
	icons, targets, _ := genCaptcha(7)
	fmt.Println(targets)
	return component.Captcha(icons, icons[targets[0]], "")
}
