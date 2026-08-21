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
func genCaptcha(rows int) ([]string, []int, error) {
	choices := make([]string, rows*rows)
	targets := []int{}
	for i := range rows * rows {
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
	icons, targets, _ := genCaptcha(3)
	fmt.Println(targets)
	return component.Captcha(icons, icons[targets[0]], "")
}
