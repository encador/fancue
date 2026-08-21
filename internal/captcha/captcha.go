package captcha

import (
	"math/rand"

	"github.com/a-h/templ"
	"github.com/encador/fancue/internal/component"
)

var emoji = []string{
	"🍆", "🥑", "🌽", "🍒", "🥜", "🍎", "🥦",
}

func genIcons(count int) []string {
	out := make([]string, count)
	for i := range count {
		out[i] = emoji[rand.Intn(len(emoji))]
	}
	return out
}

func New() templ.Component {
	n := 3
	icons := genIcons(n*n)
	return component.Captcha(icons, "")
}
