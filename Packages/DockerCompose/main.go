package DockerCompose

import (
	"AppsCompanion/Packages/Colors"
	"AppsCompanion/Packages/Figure"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func downloadAppDockerCompose() error {
	spinner := Figure.Spinner("🚀 Downloading preconfigured docker-compose.yml from web...", Colors.Blue(), "")
	spinner.Start()
	cmd := exec.Command("wget", "https://raw.githubusercontent.com/henit-chobisa/RC-Test-Environment-Companion/main/Accessories/docker-compose.yml")

	_, err := cmd.CombinedOutput()
	spinner.Stop()

	if err != nil {
		return err
	}
	fmt.Println(Colors.Yellow() + "Successfully downloaded docker-compose.yml from source")
	return nil
}

func checkFile(path string) {
	if _, err := os.Stat(path); err != nil {
		downloadAppDockerCompose()
	}
}

func Up(path string) error {
	spinner := Figure.Spinner("🚀 Starting Rocket Chat Server using Docker Compose file.", Colors.Yellow(), "")
	spinner.Start()
	time.Sleep(2 * time.Second)
	checkFile(path)
	cmd := exec.Command("docker-compose", "-f", path, "up", "-d")

	p, err := cmd.CombinedOutput()

	if err != nil || len(p) == 0 {
		spinner.Stop()
		fmt.Printf(Colors.Red()+"Docker-Compose Error : %v\n", err)
		return err
	}
	spinner.Stop()
	fmt.Printf("\n" + Colors.Yellow() + "🚀 Started Rocket Chat Server using Docker Compose file.\n")
	fmt.Println(Colors.Yellow(), fmt.Sprint("\n\n", bytes.NewBuffer(p)))
	return nil
}
