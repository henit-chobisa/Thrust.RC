package appscli

import (
	"RCTestSetup/Packages/Colors"
	"RCTestSetup/Packages/Figure"
	"bytes"
	"fmt"
	"os/exec"
)

func Install() error {
	// npm install -g
	spinner := Figure.Spinner("🔻 Installing Rocket.Chat Apps CLI globally using npm.", Colors.Cyan(), "")
	spinner.Start()
	cmd := exec.Command("npm", "install", "-g", "@rocket.chat/apps-cli")

	p, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err)
		spinner.Stop()
		return err
	}
	spinner.Stop()
	fmt.Println("\n" + Colors.Cyan() + "🔻 Installed Rocket.Chat Apps CLI globally using npm.")
	fmt.Println(Colors.Cyan(), fmt.Sprint("\n\n", bytes.NewBuffer(p)))

	return nil
}
