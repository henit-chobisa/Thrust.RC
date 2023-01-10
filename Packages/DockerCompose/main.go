package DockerCompose

import (
	"RCTestSetup/Packages/Colors"
	"RCTestSetup/Packages/Figure"
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

// func downloadAppDockerCompose() error {
// 	spinner := Figure.Spinner("🚀 Downloading preconfigured docker-compose.yml from web...", Colors.Blue(), "")
// 	spinner.Start()
// 	cmd := exec.Command("wget", "")
// }

func Up(path string) error {

	spinner := Figure.Spinner("🚀 Starting Rocket Chat Server using Docker Compose file.", Colors.Yellow(), "")
	spinner.Start()
	time.Sleep(2 * time.Second)
	cmd := exec.Command("docker-compose", "-f", path, "up", "-d")

	p, err := cmd.CombinedOutput()

	if err != nil {
		spinner.Stop()
		fmt.Printf(Colors.Red()+"Docker-Compose Error : %v", err)
		return err
	}
	spinner.Stop()
	fmt.Printf("\n" + Colors.Yellow() + "🚀 Started Rocket Chat Server using Docker Compose file.\n")
	fmt.Println(Colors.Yellow(), fmt.Sprint("\n\n", bytes.NewBuffer(p)))

	return nil
}
