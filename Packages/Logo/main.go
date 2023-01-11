package Logo

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func RocketChat() {
	fmt.Printf("\n\n\n\n")
	figure.NewColorFigure("Rocket.Chat", "banner", "Red", true).Print()
	fmt.Printf("\n")
}

func Custom(text string) {
	figure.NewColorFigure(text, "", "Red", true).Print()
}
