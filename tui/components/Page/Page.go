package Page

import (
	"RCTestSetup/tui/components"
	"RCTestSetup/tui/theme"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Task struct {
	title       string
	description string
	model       components.Model
	Output      string
}

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

type Model struct {
	tasks       list.Model
	currentTask string
	views       []string
	width       int
	height      int
	loaded      bool
}

func (m *Model) initTasks(width int, height int) {
	m.tasks = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.tasks.Title = "Current Tasks"
	m.tasks.ToggleSpinner()
	m.tasks.SetItems([]list.Item{
		Task{
			title:       "Checking Initial Configuration",
			description: "Confirm compatibility version and dependencies.",
			model:       DependencyModel{},
			Output:      "",
		},
		Task{
			title:       "Pull Containers",
			description: "Pull containers needy for running the companion, rocket.chat, companion env.",
			model:       DependencyModel{},
			Output:      "",
		},
		Task{
			title:       "Run the containers",
			description: "Pull containers needy for running the companion, rocket.chat, companion env.",
			model:       DependencyModel{},
			Output:      "",
		},
		Task{
			title:       "Setup Project Environment",
			description: "Pull containers needy for running the companion, rocket.chat, companion env.",
			model:       DependencyModel{},
			Output:      "",
		},
		Task{
			title:       "Pull Companion Containers",
			description: "Pull containers needy for running the companion, rocket.chat, companion env.",
			model:       DependencyModel{},
			Output:      "",
		},
		Task{
			title:       "Companion Logs",
			description: "Pull containers needy for running the companion, rocket.chat, companion env.",
			model:       DependencyModel{},
			Output:      "",
		},
	})

}

func New() *Model {
	return &Model{
		currentTask: "Checking Initial Configuration",
		views:       []string{},
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height, m.width = msg.Height, msg.Width
		if !m.loaded {
			m.loaded = true
		}
		m.initTasks(m.Width()/2, m.Height())
	}
	var cmd tea.Cmd
	m.tasks, cmd = m.tasks.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	if !m.loaded {
		return "Loading..."
	}
	pageStyle := lipgloss.NewStyle().Height(m.height).Width(m.width/2).MarginLeft(15).Border(lipgloss.RoundedBorder()).BorderForeground(theme.PrimaryColour).Padding(2, 3, 1, 3)
	return lipgloss.JoinHorizontal(lipgloss.Center, m.tasks.View(), pageStyle.Render(lipgloss.JoinVertical(lipgloss.Left)))
}

func (m *Model) Resize(width, height int) components.Model {
	m.height = height
	m.width = width
	if !m.loaded {
		m.loaded = true
	}
	m.initTasks(m.Width()/2, m.Height())
	return m
}

func (m *Model) Width() int {
	return m.width
}

func (m *Model) Height() int {
	return lipgloss.Height(m.View())
}
