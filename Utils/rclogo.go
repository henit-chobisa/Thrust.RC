package Utils

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func PrintRCLogo() {
	figure.NewColorFigure("Rocket.Chat", "", "Red", true).Print()
	fmt.Printf("\n\n")
}
