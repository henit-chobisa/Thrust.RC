package footer

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	err    error
	task   string
	width  int
	styles *Styles
}

func New(err error, task string) Model {
	return Model{
		err:    err,
		task:   task,
		styles: DefaultStyles(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	if m.err != nil {
		return m.styles.err.Render(m.err.Error())
	} else {
		return m.styles.status.Render(m.task)
	}
}

func (m Model) Resize(width, height int) Model {
	m.width = width
	return m
}

func (m Model) Width() int {
	return m.width
}

func (m Model) Height() int {
	return lipgloss.Height(m.View())
}
