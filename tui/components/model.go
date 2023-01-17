package components

import tea "github.com/charmbracelet/bubbletea"

type Model interface {
	tea.Model
	Resize(width, height int) Model
	Width() int
	Height() int
}
