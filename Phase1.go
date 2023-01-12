package main

import (
	appscli "AppsCompanion/Packages/AppsCli"
	"AppsCompanion/Packages/Colors"
	"AppsCompanion/Packages/ConfigReader"
	constants "AppsCompanion/Packages/Constants"
	"AppsCompanion/Packages/DockerCompose"
	"AppsCompanion/Packages/Figure"
	"AppsCompanion/Packages/Logo"
	"fmt"
)

func showInfo(name string) {
	Logo.RocketChat()
	Logo.Custom(name)
	fmt.Printf("\n\n\n")
	fmt.Println(Colors.Blue() + "Phase 1 : Intiating Rocket Chat Apps Test Environment\n" + Figure.Line())
}

func InitiatePhase1(data map[string]interface{}, appDir string) {

	appData := ConfigReader.ReadConfig(fmt.Sprintf("%v/app.json", appDir))

	showInfo(fmt.Sprintf("%v App", appData["name"]))

	composePath := constants.DockerCompose_default

	if data["composeFilePath"] != nil {
		composePath = fmt.Sprintf("%v", data["composeFilePath"])
	}

	DockerCompose.Up(composePath)

	appscli.Install()

	fmt.Printf("\n")
}
