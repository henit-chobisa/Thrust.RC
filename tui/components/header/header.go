package header

import (
	"strings"

	"RCTestSetup/tui/components"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	name        string
	version     string
	description string
	width       int
	Styles      *Styles
}

func New(name, version, description string) Model {
	return Model{
		name:        name,
		description: description,
		version:     version,
		Styles:      DefaultStyles(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var b strings.Builder

	nameVersion := lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.Styles.Name.Padding(0, 2).Render("Rocket.Chat"),
		m.Styles.Version.Padding(0, 2).Render(m.version),
	)

	desc := m.Styles.Description.Render(m.description)

	banner := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().MarginBottom(1).Render(nameVersion),
		desc,
	)

	b.WriteString(m.Styles.Border.
		MarginBottom(1).
		Width(m.width).
		Render(banner))

	return b.String()
}

func (m Model) Resize(width, height int) components.Model {
	m.width = width
	return m
}

func (m Model) Width() int {
	return m.width
}

func (m Model) Height() int {
	return lipgloss.Height(m.View())
}
