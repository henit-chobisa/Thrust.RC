package Page

import (
	"thrust/tui/theme"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type PageStyle list.Styles

func DefaultStyles() *PageStyle {
	s := &PageStyle{}
	s.Title = lipgloss.NewStyle().Background(theme.SecondaryColour).Margin(1, 0, 1, 0).Padding(0, 1, 0, 1).Bold(true)
	s.StatusBar = lipgloss.NewStyle().Foreground(theme.FeintColour).Margin(0, 0, 1, 0)

	return s
}

type DefaultItemStyles struct {
	NormalTitle lipgloss.Style
	NormalDesc  lipgloss.Style

	// The selected item state.
	SelectedTitle lipgloss.Style
	SelectedDesc  lipgloss.Style

	// The dimmed state, for when the filter input is initially activated.
	DimmedTitle lipgloss.Style
	DimmedDesc  lipgloss.Style

	// Charcters matching the current filter, if any.
	FilterMatch lipgloss.Style
}

func NewDefaultItemStyles() list.DefaultItemStyles {

	s := list.NewDefaultItemStyles()

	s.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: string(theme.SecondaryColour), Dark: string(theme.SecondaryColour)}).
		Foreground(lipgloss.AdaptiveColor{Light: string(theme.SecondaryColour), Dark: string(theme.SecondaryColour)}).
		Padding(0, 0, 0, 1)

	s.SelectedDesc = s.SelectedTitle.Copy().
		Foreground(lipgloss.AdaptiveColor{Light: string(theme.FeintColour), Dark: string(theme.FeintColour)})

	return s
}
