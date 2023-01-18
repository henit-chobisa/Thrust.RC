package components

import tea "github.com/charmbracelet/bubbletea"

type Model[T any] interface {
	tea.Model
	New() T
	Resize(width, height int) T
	Width() int
	Height() int
}
