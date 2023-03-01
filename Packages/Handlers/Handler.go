package Handlers

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	constants "thrust/Packages/Constants"
	"thrust/Packages/DockerSDK"
	"thrust/Packages/DockerSDK/DefaultContainers"
	initiateadmin "thrust/Packages/InitiateAdmin"
	models "thrust/Packages/Models"
	"thrust/Utils"
	"thrust/tui/components/Page"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/spf13/viper"
)

func HandleDependencyCheck() error {
	// confirming initial configuration dependency
	fmt.Printf(constants.Blue + "\n🐳 Checking Initial dependencies required for running the companion\n\n")
	dependencyModel := Page.NewDependencyModel()
	fmt.Println(dependencyModel.View())

	if dependencyModel.Err != nil {
		return dependencyModel.Err
	}
	return nil
}

func HandlePullingImages() (map[string]string, error) {
	// first let's check for the images that are currently present in the system

	isAppMode := viper.GetBool("isAppMode")

	fmt.Printf(constants.Blue + "🐳 Verifying Required Docker Images for Companion\n\n" + constants.White)

	client, err := DockerSDK.GetNewClient()
	if err != nil {
		fmt.Println(constants.Red + "× Something went wrong, error accessing the client")
		return nil, err
	}

	filter := filters.NewArgs()

	filter.Add("reference", constants.RocketChatImage)
	filter.Add("reference", constants.MongoDBImage)

	if isAppMode {
		filter.Add("reference", constants.CompanionImage)
	}

	imageSummary, err := client.FindImages(filter)

	if err != nil {
		fmt.Println(constants.Red + "× Something went wrong, error searching for images")
		return nil, err
	}

	imagesToPull := map[string]string{constants.RocketChatImage: "Rocket.Chat", constants.MongoDBImage: "MongoDB"}

	if isAppMode {
		imagesToPull[constants.CompanionImage] = "Apps.Companion"
	}

	for _, image := range *imageSummary {
		fmt.Println(Utils.Tick() + imagesToPull[image.RepoTags[0]] + " : " + image.RepoTags[0])
		delete(imagesToPull, image.RepoTags[0])
	}

	for key, values := range imagesToPull {
		fmt.Println(Utils.Cross() + values + " : " + key)
	}

	client.Client.Close()

	return imagesToPull, nil
}

func PullImages(images map[string]string) error {
	fmt.Printf(constants.Blue + "\n🐳 Pulling Required Images\n" + constants.White)

	var tasks sync.WaitGroup
	tasks.Add(len(images))

	client, err := DockerSDK.GetNewClient()

	if err != nil {
		return err
	}

	for key, value := range images {
		fmt.Printf("\nStarted Pulling %v image, logs would be displayed before you\n", value)
		go client.PullImage(key, &tasks)
	}

	tasks.Wait()

	fmt.Println()

	client.Client.Close()

	return nil
}

func getRunningContainers(client *DockerSDK.Docker) (*[]types.Container, error) {
	filter := filters.NewArgs()

	filter.Add("ancestor", constants.RocketChatImage)
	filter.Add("ancestor", constants.MongoDBImage)
	if viper.GetBool("appMode") {
		filter.Add("ancestor", constants.CompanionImage)
	}
	filter.Add("status", "running")
	filter.Add("network", "RCAPPSDEFAULT")

	return client.FindContainers(filter)
}

func CheckRequiredContainers(app *models.AppInfo) (containersToStart map[string]string, companionStart bool, containerIDs map[string]string, err error) {
	fmt.Printf(constants.Blue + "\n🐳 Finding Running Containers\n\n" + constants.White)
	client, err := DockerSDK.GetNewClient()
	containerIDs = map[string]string{}
	isAppMode := viper.GetBool("appMode")

	if err != nil {
		return nil, false, containerIDs, err
	}

	if err != nil {
		return nil, false, containerIDs, err
	}

	containers, err := getRunningContainers(client)

	if err != nil {
		return nil, false, containerIDs, err
	}

	containersToStart = map[string]string{constants.RocketChatImage: "Rocket.Chat", constants.MongoDBImage: "MongoDB"}

	companionName := ""

	if isAppMode {
		containersToStart[constants.CompanionImage] = "Apps.Companion"
		companionName = "/companion_" + app.Name + app.Id + app.Version
	}

	companionStart = false

	for _, container := range *containers {
		if companionName != container.Names[0] && !strings.HasPrefix(container.Names[0], "/companion") {
			fmt.Println(Utils.Tick() + containersToStart[container.Image] + " : " + container.ID + " : " + container.Names[0] + " : " + container.Image + " : " + container.Status)

			containerIDs[container.Image] = container.ID

			delete(containersToStart, container.Image)
			continue
		} else if companionName == container.Names[0] {
			fmt.Println(Utils.Tick() + containersToStart[container.Image] + " : " + container.ID + " : " + container.Names[0] + " : " + container.Image + " : " + container.Status)

			if container.Image == constants.CompanionImage {
				containerIDs[constants.CompanionImage] = container.ID
			}

			if len(containersToStart) > 0 {
				delete(containersToStart, container.Image)
			}
		}
	}

	for key, value := range containersToStart {
		if value == "Apps.Companion" {
			companionStart = true
		}
		fmt.Println(Utils.Cross() + value + " : " + key)
	}

	return containersToStart, companionStart, containerIDs, nil
}

func createDefaultNetwork() (id string, err error) {

	client, err := DockerSDK.GetNewClient()

	if err != nil {
		return "", nil
	}

	id, err = client.CreateNetwork()
	return id, err
}

func StartContainersWithDefaultNetwork(containers map[string]string) (containerIDs map[string]string, err error) {

	fmt.Printf(constants.Blue + "\n🐳 Starting Required Containers for Companion\n\n" + constants.White)

	containerIDs = map[string]string{constants.RocketChatImage: "Rocket.Chat", constants.MongoDBImage: "MongoDB"}

	defaultNetworkId, err := createDefaultNetwork()

	if err != nil {
		return containerIDs, err
	}

	client, err := DockerSDK.GetNewClient()

	if err != nil {
		return containerIDs, err
	}
	if _, ok := containers[constants.MongoDBImage]; ok {
		mongodbID, err := DefaultContainers.LaunchMongoDbContainer(*client, defaultNetworkId)

		if err != nil {
			fmt.Println(Utils.Cross() + "Error Starting Mongodb Container for Rocket.Chat, Aborting Operation")
			return containerIDs, err
		}

		containerIDs[constants.MongoDBImage] = mongodbID

		fmt.Println(Utils.Tick() + "Started Mongodb Container for Rocket.Chat with ID " + mongodbID)
	}

	if _, ok := containers[constants.RocketChatImage]; ok {
		if !viper.GetBool("appMode") {
			fmt.Println(Utils.Tick() + "Starting Rocket.Chat Container for assitance")
		}
		rocketChatID, err := DefaultContainers.LaunchRocketChatContainer(*client, defaultNetworkId)

		if err != nil {
			fmt.Println(Utils.Cross() + "Error Starting Rocket.Chat Container for Rocket.Chat, Aborting Operation")
			return containerIDs, err
		}
		fmt.Println(rocketChatID + " Rocket.Chat ID")
		containerIDs[constants.RocketChatImage] = rocketChatID

		if viper.GetBool("appMode") {
			fmt.Println(Utils.Tick() + "Started Rocket.Chat Container with ID " + rocketChatID)
		}

	}

	return containerIDs, nil
}

func CreateAdminUser() error {

	fmt.Printf(constants.Blue + "\n🐳 Creating Admin User for Rocket.Chat Instance\n\n" + constants.White)

	initiateadmin.Initiate()

	return nil
}

func StartCompanionContainer(appDir string, app *models.AppInfo) error {

	fmt.Printf(constants.Blue + "\n🐳 Starting Required Containers for Companion\n\n" + constants.White)

	defaultNetworkId, err := createDefaultNetwork()

	if err != nil {
		return err
	}

	client, err := DockerSDK.GetNewClient()

	if err != nil {
		return err
	}

	companionName := "companion_" + app.Name + app.Id + app.Version

	_, err = DefaultContainers.LaunchCompanionContainer(*client, defaultNetworkId, appDir, companionName)

	if err != nil {
		return err
	}

	return nil
}

func ShowLogs(containerID string) error {
	client, _ := DockerSDK.GetNewClient()

	out, err := client.Client.ContainerLogs(context.Background(), containerID, types.ContainerLogsOptions{ShowStdout: true, Follow: true})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, out)
	return nil
}

func Cleanup(app *models.AppInfo) error {

	client, err := DockerSDK.GetNewClient()

	if err != nil {
		return err
	}

	runningContainers, err := getRunningContainers(client)

	var companionName string = ""

	if viper.GetBool("appMode") {
		companionName = "/companion_" + app.Name + app.Id + app.Version
	}

	if err != nil {
		return err
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(*runningContainers))
	fmt.Println(len(*runningContainers))
	for _, container := range *runningContainers {
		if viper.GetBool("appMode") {
			if len(*runningContainers) == 3 || container.Names[0] == companionName {
				go client.RemoveContainer(wg, container.ID)
			}
		} else {
			go client.RemoveContainer(wg, container.ID)
		}
	}

	wg.Wait()

	return nil
}
