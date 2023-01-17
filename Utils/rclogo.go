package Utils

import (
	"github.com/common-nighthawk/go-figure"
)

func PrintRCLogo() string {
	return figure.NewFigure("Rocket.Chat", "", true).String()
}
