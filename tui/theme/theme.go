package theme

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	PrimaryColour        = lipgloss.Color("#F5455C")
	SecondaryColour      = lipgloss.Color("#1D74F5")
	BorderColour         = lipgloss.Color("#F7F8FA")
	FeintColour          = lipgloss.Color("#807d8a")
	VeryFeintColour      = lipgloss.Color("#5e5e5e")
	TextColour           = lipgloss.Color("#F7F8FA")
	HighlightColour      = lipgloss.Color("#bf31f7")
	HighlightFeintColour = lipgloss.Color("#b769d6")
	AmberColour          = lipgloss.Color("#e68a35")
	GreenColour          = lipgloss.Color("#26a621")
	WhiteColour          = lipgloss.Color("#F7F8FA")

	AppStyle            = lipgloss.NewStyle().Margin(1)
	TextStyle           = lipgloss.NewStyle().Foreground(TextColour)
	FeintTextStyle      = lipgloss.NewStyle().Foreground(FeintColour)
	VeryFeintTextStyle  = lipgloss.NewStyle().Foreground(VeryFeintColour)
	HightlightTextStyle = lipgloss.NewStyle().Foreground(HighlightColour)
)
