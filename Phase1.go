package main

// import (
// 	appscli "RCTestSetup/Packages/AppsCli"
// 	"RCTestSetup/Packages/ConfigReader"
// 	constants "RCTestSetup/Packages/Constants"
// 	"RCTestSetup/Packages/DockerCompose"
// 	"RCTestSetup/Packages/Logo"
// 	"fmt"
// )

// func showInfo(name string) {
// 	Logo.RocketChat()
// 	Logo.Custom(name)
// 	fmt.Printf("\n\n\n")
// 	fmt.Println(constants.Blue + "Phase 1 : Intiating Rocket Chat Apps Test Environment\n" + constants.Line)
// }

// func InitiatePhase1(data map[string]interface{}, appDir string) {

// 	appData := ConfigReader.ReadConfig(fmt.Sprintf("%v/app.json", appDir))

// 	showInfo(fmt.Sprintf("%v App", appData["name"]))

// 	composePath := constants.DockerCompose_default

// 	if data["composeFilePath"] != nil {
// 		composePath = fmt.Sprintf("%v", data["composeFilePath"])
// 	}

// 	DockerCompose.Up(composePath)

// 	appscli.Install()

// 	fmt.Printf("\n")
// }
