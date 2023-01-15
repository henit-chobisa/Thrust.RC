package main

// import (
// 	constants "RCTestSetup/Packages/Constants"
// 	"RCTestSetup/Packages/InstallApp"
// 	"RCTestSetup/Packages/Logo"
// 	"fmt"
// )

// func showClosingInfo() {
// 	Logo.RocketChat()

// 	fmt.Printf("\n")
// 	fmt.Println(constants.Green + "🚀 Testing environment ready for using your app\n" + constants.Line)
// 	fmt.Println("\n" + constants.Green + "Now you can open http://localhost:3000, use the credentials given in the config.json file and test the app.\nIf you are using web version of Gitpod make sure to install Gitpod's Local Companion.\n\nAuthor: Henit Chobisa(@henit-chobisa)\n✨ Make sure to follow Rocket.Chat and me...")
// }

// func InitiatePhase3(data map[string]interface{}, appDir string) {

// 	fmt.Printf("\n\n")
// 	fmt.Println(constants.Blue + "Phase 3 : Installing App into Rocket.Chat Server\n" + constants.Line)

// 	if data["admin"] == nil {
// 		InstallApp.Install(appDir, "http://localhost:3000", "user0", "123456", false)
// 	} else {
// 		user := data["admin"].(map[string]interface{})
// 		InstallApp.Install(appDir, "http://localhost:3000", fmt.Sprintf("%v", user["username"]), fmt.Sprintf("%v", user["pass"]), true)
// 	}

// 	showClosingInfo()
// }
