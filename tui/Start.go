package tui

import (
	"RCTestSetup/Utils"
	"RCTestSetup/enums"

	tea "github.com/charmbracelet/bubbletea"
)

type StartModel struct {
	config          string
	watcher         bool
	username        string
	email           string
	password        string
	name            string
	virtual         bool
	deps            bool
	composeFilePath string
}

func InitialModel() StartModel {
	return StartModel{
		config:          Utils.GetConfig(enums.Config).(string),
		watcher:         Utils.GetConfig(enums.Watcher).(bool),
		username:        Utils.GetConfig(enums.Username).(string),
		email:           Utils.GetConfig(enums.Email).(string),
		password:        Utils.GetConfig(enums.Password).(string),
		name:            Utils.GetConfig(enums.Name).(string),
		virtual:         Utils.GetConfig(enums.Virtual).(bool),
		deps:            Utils.GetConfig(enums.Deps).(bool),
		composeFilePath: Utils.GetConfig(enums.ComposeFilePath).(string),
	}
}

func (s StartModel) Init() tea.Cmd {
	Utils.PrintRCLogo()
	return nil
}

func (s StartModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	tea.Quit()
	return s, nil
}

func (s StartModel) View() string {

	return "hello, the start program is up and running..."
}
