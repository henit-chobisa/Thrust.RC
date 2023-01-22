package Page

import (
	"AppsCompanion/enums"
	"AppsCompanion/tui/theme"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Task struct {
	title       string
	description string
	model       PageModel
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
	currentTask enums.Task
	views       []string
	width       int
	height      int
	loaded      bool
}

type LoadingStart int

func startLoading() tea.Msg {
	return LoadingStart(1)
}

// type LoadingStop int

// func stopLoading() tea.Msg {
// 	return LoadingStop(1)
// }

// func (m *Model) ExecuteList() tea.Msg {
// 	return LoadingStart(1)
// }

func (m *Model) initTasks(width int, height int) {

	delegate := list.NewDefaultDelegate()
	delegate.Styles = NewDefaultItemStyles()

	m.tasks = list.New([]list.Item{}, delegate, width, height)
	m.tasks.Title = "Current Tasks"
	m.tasks.Styles = list.Styles(*DefaultStyles())
	m.tasks.SetShowStatusBar(false)
	m.tasks.SetShowHelp(false)

	m.tasks.SetItems([]list.Item{
		Task{
			title:       enums.Check_Initial_Configuration.String(),
			description: "Confirm compatibility version and dependencies.",
			model:       NewDependencyModel(),
			Output:      "",
		},
		Task{
			title:       enums.Pull_Containers.String(),
			description: "Pull containers needy for running the companion, rocket.chat, companion env.",
			model:       NewDependencyModel(),
			Output:      "",
		},
		Task{
			title:       enums.Run_containers.String(),
			description: "Pull containers needy for running the companion, rocket.chat, companion env.",
			model:       NewDependencyModel(),
			Output:      "",
		},
		Task{
			title:       enums.Setup_Project_Environment.String(),
			description: "Pull containers needy for running the companion, rocket.chat, companion env.",
			model:       NewDependencyModel(),
			Output:      "",
		},
		Task{
			title:       enums.Setup_Project_Environment.String(),
			description: "Pull containers needy for running the companion, rocket.chat, companion env.",
			model:       NewDependencyModel(),
			Output:      "",
		},
		Task{
			title:       enums.Show_Companion_Logs.String(),
			description: "Pull containers needy for running the companion, rocket.chat, companion env.",
			model:       NewDependencyModel(),
			Output:      "",
		},
	})
}

func New() *Model {
	return &Model{
		currentTask: 1,
		views:       []string{},
	}
}

func (m *Model) Init() tea.Cmd {
	return tea.Sequence()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height, m.width = msg.Height, msg.Width
		// case loadingStart:
		// 	switch msg {
		// 	case 1:
		// 		if m.loaded {
		// 			m.tasks.Select(3)
		// 		}
		// 	}
		// }
	}
	if m.loaded {
		var cmd tea.Cmd
		m.tasks, cmd = m.tasks.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m *Model) View() string {
	if !m.loaded {
		return "Loading..."
	}
	pageStyle := lipgloss.NewStyle().Height(m.height).Width(m.width/2).MarginLeft(15).Border(lipgloss.RoundedBorder()).BorderForeground(theme.PrimaryColour).Padding(2, 3, 1, 3)
	return lipgloss.JoinHorizontal(lipgloss.Center, m.tasks.View(), pageStyle.Render(lipgloss.JoinVertical(lipgloss.Left, m.tasks.SelectedItem().(Task).model.View())))
}

func (m *Model) Resize(width, height int) *Model {
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
