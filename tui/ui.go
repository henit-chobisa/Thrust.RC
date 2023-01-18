package tui

import (
	"RCTestSetup/tui/components/Page"
	"RCTestSetup/tui/components/footer"
	"RCTestSetup/tui/components/header"
	"RCTestSetup/tui/theme"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type UI struct {
	header header.Model
	footer footer.Model
	page   Page.Model
}

func New() UI {
	return UI{
		header: header.New("Rocket.Chat", "1.0.0", "Companion for Rocket.Chat Apps"),
		footer: footer.New(nil, "Hello we are going to initiate the task right now"),
		page:   *Page.New(),
	}
}

func (u UI) Init() tea.Cmd {
	return nil
}

func (u UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		x, y := u.margins()

		u.header = u.header.Resize(msg.Width-x, u.header.Height())
		u.footer = u.footer.Resize(msg.Width-x, u.footer.Height())

		pageX := msg.Width - x
		pageY := msg.Height - (y + u.header.Height() + u.footer.Height())

		u.page = *(u.page).Resize(pageX, pageY)

		return u, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+q":
			return u, tea.Quit
		}
	}
	u.page.Update(msg)
	return u, nil
}

func (u UI) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, u.header.View(), u.page.View(), u.footer.View())
}

func (u UI) margins() (int, int) {
	s := theme.AppStyle.Copy()
	return s.GetHorizontalFrameSize(), s.GetVerticalFrameSize()
}
