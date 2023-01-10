package InstallApp

import (
	"RCTestSetup/Packages/Colors"
	"RCTestSetup/Packages/Figure"
	"bytes"
	
	"fmt"
	"os/exec"
)

func Install(path string, url string, username string, password string) error {
	spinner := Figure.Spinner(" Installing App into the Rocket.Chat testing server", Colors.Purple(), "")
	spinner.Start()
	cmd := exec.Command("rc-apps", "deploy", "--url", url, "--username", username, "--password", password)
	cmd.Dir = path

	p, err := cmd.CombinedOutput()

	if err != nil {
		spinner.Stop()
		fmt.Println(err)
		return err
	}
	spinner.Stop()

	fmt.Printf("\n" + Colors.Purple() + "🚀 Successfully installed app into the Rocket.Chat testing server\n")
	fmt.Println(Colors.Purple(), fmt.Sprint("\n\n", bytes.NewBuffer(p)))
	return nil
}
