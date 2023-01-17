package Page

import (
	"RCTestSetup/Packages/DockerSDK"
	"RCTestSetup/tui/components"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	checkMarkTrue  = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	checkMarkFalse = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff2255")).SetString("×")
)

type DependencyModel struct {
	dockerEnabled    bool
	dockerVersion    string
	dockerAPIVersion string
	dockerEngineType string
	os               string
	err              error
	width            int
	height           int
}

func (d DependencyModel) New() DependencyModel {
	version, err := DockerSDK.GetVersionInfo()
	if err != nil {
		return DependencyModel{
			err: err,
		}
	}
	return DependencyModel{
		dockerEnabled:    true,
		dockerAPIVersion: version.APIVersion,
		dockerVersion:    version.Version,
		dockerEngineType: version.Platform.Name,
		os:               version.Arch,
		err:              nil,
	}
}

func (d DependencyModel) Init() tea.Cmd {
	return nil
}

func (d DependencyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return d, tea.Quit
		case "ctrl+r":
			d = d.New()
		}
		return d, tea.ClearScreen
	}
	return d, nil
}

func (d DependencyModel) View() string {

	if d.err != nil {
		info := lipgloss.NewStyle().Render("Something is probably wrong, confirm that your docker-daemon is running, please.")
		return fmt.Sprintf("%s %s\n\nPress (q) to quit\nPress (ctrl+r) to reload", checkMarkFalse, info)
	}

	textStyle := lipgloss.NewStyle().Margin(0)

	return lipgloss.JoinVertical(lipgloss.Left, textStyle.Render(fmt.Sprintf("%s Compatible Docker Version Found: %s\n", checkMarkTrue, d.dockerVersion)), textStyle.Render(fmt.Sprintf("%s Compatible Docker Engine Type : %s\n", checkMarkTrue, d.dockerEngineType)), textStyle.Render(fmt.Sprintf("%s Running Docker API Version: %s\n", checkMarkTrue, d.dockerAPIVersion)), textStyle.Render(fmt.Sprintf("%s On Operating System: %s\n", checkMarkTrue, d.os)))
}

func (d DependencyModel) Resize(width, height int) components.Model {
	d.width = width
	return d
}

func (d DependencyModel) Width() int {
	return d.width
}

func (d DependencyModel) Height() int {
	return lipgloss.Height(d.View())
}
