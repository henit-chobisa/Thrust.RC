package InstallApp

import (
	constants "RCTestSetup/Packages/Constants"
	"RCTestSetup/Packages/UIAssets"
	"bytes"

	"fmt"
	"os/exec"
)

func Install(path string, url string, username string, password string, reload bool) error {
	var message string
	if reload {
		message = " Uploading the updated app into the server..."
	} else {
		message = " Installing App into the Rocket.Chat testing server"
	}
	spinner := UIAssets.Spinner(message, constants.Purple, "")
	spinner.Start()
	cmd := exec.Command("rc-apps", "deploy", "--url", url, "--username", username, "--password", password)
	cmd.Dir = path

	p, err := cmd.CombinedOutput()

	if err != nil {
		spinner.Stop()
		fmt.Println(constants.Red + err.Error())
		return err
	}
	spinner.Stop()

	fmt.Printf("\n" + constants.Purple + "🚀 Successfully installed app into the Rocket.Chat testing server\n")
	fmt.Println(constants.Purple, fmt.Sprint("\n\n", bytes.NewBuffer(p)))
	return nil
}
