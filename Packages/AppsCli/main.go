package appscli

import (
	constants "RCTestSetup/Packages/Constants"
	"RCTestSetup/Packages/UIAssets"
	"bytes"
	"fmt"
	"os/exec"
)

func Install() error {
	spinner := UIAssets.Spinner("🔻 Installing Rocket.Chat Apps CLI globally using npm.", constants.Cyan, "")
	spinner.Start()
	cmd := exec.Command("sudo", "npm", "install", "-g", "@rocket.chat/apps-cli")

	p, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		spinner.Stop()
		return err
	}
	spinner.Stop()
	fmt.Println("\n" + constants.Cyan + "🔻 Installed Rocket.Chat Apps CLI globally using npm.")
	fmt.Println(constants.Cyan, fmt.Sprint("\n\n", bytes.NewBuffer(p)))

	return nil
}
