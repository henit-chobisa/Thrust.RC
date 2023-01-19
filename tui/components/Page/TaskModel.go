package Page

import tea "github.com/charmbracelet/bubbletea"

type PageModel interface {
	tea.Model
	Run() tea.Msg
	Width() int
	Height() int
	Resize(width, height int)
	// New(pageModel *PageModel)
}
