package DockerCompose

import (
	constants "RCTestSetup/Packages/Constants"
	"RCTestSetup/Packages/UIAssets"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func downloadAppDockerCompose() error {
	spinner := UIAssets.Spinner("🚀 Downloading preconfigured docker-compose.yml from web...", constants.Blue, "")
	spinner.Start()
	cmd := exec.Command("wget", "https://raw.githubusercontent.com/henit-chobisa/RC-Test-Environment-Companion/main/Accessories/docker-compose.yml")

	_, err := cmd.CombinedOutput()
	spinner.Stop()

	if err != nil {
		return err
	}
	fmt.Println(constants.Yellow + "Successfully downloaded docker-compose.yml from source")
	return nil
}

func checkFile(path string) {
	if _, err := os.Stat(path); err != nil {
		downloadAppDockerCompose()
	}
}

func Up(path string) error {
	spinner := UIAssets.Spinner("🚀 Starting Rocket Chat Server using Docker Compose file.", constants.Yellow, "")
	spinner.Start()
	time.Sleep(2 * time.Second)
	checkFile(path)
	cmd := exec.Command("docker-compose", "-f", path, "up", "-d")

	p, err := cmd.CombinedOutput()

	if err != nil || len(p) == 0 {
		spinner.Stop()
		fmt.Printf(constants.Red+"Docker-Compose Error : %v\n", err)
		return err
	}
	spinner.Stop()
	fmt.Printf("\n" + constants.Yellow + "🚀 Started Rocket Chat Server using Docker Compose file.\n")
	fmt.Println(constants.Yellow, fmt.Sprint("\n\n", bytes.NewBuffer(p)))
	return nil
}
