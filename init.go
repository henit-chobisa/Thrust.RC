package main

import (
	appscli "RCTestSetup/Packages/AppsCli"
	"RCTestSetup/Packages/Colors"
	"RCTestSetup/Packages/ConfigReader"
	"RCTestSetup/Packages/DockerCompose"
	"RCTestSetup/Packages/Figure"
	initiateadmin "RCTestSetup/Packages/InitiateAdmin"
	"RCTestSetup/Packages/InstallApp"
	"RCTestSetup/Packages/Logo"
	"fmt"
	"os/exec"
)

func main() {
	data := ConfigReader.ReadConfig("config.json")
	appDir := "./"
	if data["appDir"] != nil {
		appDir = fmt.Sprintf("%v", data["appDir"])
	}
	appData := ConfigReader.ReadConfig(fmt.Sprintf("%v/app.json", appDir))

	Logo.RocketChat()
	Logo.Custom(fmt.Sprintf("%v App", appData["name"]))
	fmt.Printf("\n\n\n")
	fmt.Println(Colors.Blue() + "Phase 1 : Intiating Rocket Chat Apps Test Environment\n" + Figure.Line())

	composePath := "./docker-compose.yml"
	if data["composeFilePath"] != nil {
		composePath = fmt.Sprintf("%v", data["composeFilePath"])
	}
	DockerCompose.Up(composePath)
	appscli.Install()

	fmt.Printf("\n")
	fmt.Println(Colors.Blue() + "Phase 2 : Configuring Rocket.Chat App, installing admin\n" + Figure.Line())
	initiateadmin.Initiate(data)

	fmt.Printf("\n\n")
	fmt.Println(Colors.Blue() + "Phase 3 : Installing App into Rocket.Chat Server\n" + Figure.Line())

	user := data["admin"].(map[string]interface{})

	InstallApp.Install(fmt.Sprintf("%v", data["appDir"]), "http://localhost:3000", fmt.Sprintf("%v", user["username"]), fmt.Sprintf("%v", user["pass"]))

	Logo.RocketChat()

	fmt.Printf("\n")
	fmt.Println(Colors.Green() + "🚀 Testing environment ready for using your app\n" + Figure.Line())
	fmt.Println("\n" + Colors.Green() + "Now you can open http://localhost:3000, use the credentials given in the config.json file and test the app.\nIf you are using web version of Gitpod make sure to install Gitpod's Local Companion.\n\nAuthor: Henit Chobisa(@henit-chobisa)\n✨ Make sure to follow Rocket.Chat and me...\n\n")

	exec.Command("gp", "preview", "http://localhost:3000", "--external").Output()
}
