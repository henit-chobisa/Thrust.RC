package footer

import (
	"thrust/tui/theme"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	err    lipgloss.Style
	status lipgloss.Style
}

func DefaultStyles() *Styles {
	s := &Styles{}
	s.status = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Background(theme.SecondaryColour).Padding(0, 2, 0, 2).Bold(true).Width(200)
	s.err = lipgloss.NewStyle().Foreground(theme.WhiteColour).Background(theme.PrimaryColour).Padding(0, 2, 0, 2).Bold(true).Width(lipgloss.NewStyle().GetHorizontalFrameSize())
	return s
}
