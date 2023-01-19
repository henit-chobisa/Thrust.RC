package Utils

import (
	constants "RCTestSetup/Packages/Constants"

	"github.com/common-nighthawk/go-figure"
)

func PrintRCLogo() string {
	return figure.NewFigure("Rocket.Chat", "", true).String()
}

func Tick() string {
	return constants.Green + "✓ " + constants.White
}

func Cross() string {
	return constants.Red + "× " + constants.White
}
