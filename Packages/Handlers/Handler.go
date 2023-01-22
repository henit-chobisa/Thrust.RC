package Handlers

import (
	constants "RCTestSetup/Packages/Constants"
	"RCTestSetup/Packages/DockerSDK"
	"RCTestSetup/Packages/DockerSDK/DefaultContainers"
	initiateadmin "RCTestSetup/Packages/InitiateAdmin"
	"RCTestSetup/Utils"
	"RCTestSetup/tui/components/Page"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
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

	fmt.Printf(constants.Blue + "🐳 Verifying Required Docker Images for Companion\n\n" + constants.White)

	client, err := DockerSDK.GetNewClient()
	if err != nil {
		fmt.Println(constants.Red + "× Something went wrong, error accessing the client")
		return nil, err
	}

	filter := filters.NewArgs()

	filter.Add("reference", constants.RocketChatImage)
	filter.Add("reference", constants.MongoDBImage)
	filter.Add("reference", constants.CompanionImage)

	imageSummary, err := client.FindImages(filter)

	if err != nil {
		fmt.Println(constants.Red + "× Something went wrong, error searching for images")
		return nil, err
	}

	imagesToPull := map[string]string{constants.RocketChatImage: "Rocket.Chat", constants.MongoDBImage: "MongoDB", constants.CompanionImage: "Apps.Companion"}

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

func CheckRequiredContainers() (containersToStart map[string]string, companionStart bool, companionID string, err error) {
	fmt.Printf(constants.Blue + "\n🐳 Finding Running Containers\n\n" + constants.White)
	client, err := DockerSDK.GetNewClient()
	if err != nil {
		return nil, false, "", err
	}

	filter := filters.NewArgs()

	filter.Add("ancestor", constants.RocketChatImage)
	filter.Add("ancestor", constants.MongoDBImage)
	filter.Add("ancestor", constants.CompanionImage)
	filter.Add("status", "running")
	filter.Add("network", "RCAPPSDEFAULT")

	containers, err := client.FindContainers(filter)

	if err != nil {
		return nil, false, "", err
	}

	containersToStart = map[string]string{constants.RocketChatImage: "Rocket.Chat", constants.MongoDBImage: "MongoDB", constants.CompanionImage: "Apps.Companion"}

	companionStart = false

	for _, container := range *containers {
		fmt.Println(Utils.Tick() + containersToStart[container.Image] + " : " + container.ID + " : " + container.Names[0] + " : " + container.Image + " : " + container.Status)
		if container.Image == constants.CompanionImage {
			companionID = container.ID
		}
		delete(containersToStart, container.Image)
	}

	for key, value := range containersToStart {
		if value == "Apps.Companion" {
			companionStart = true
		}
		fmt.Println(Utils.Cross() + value + " : " + key)
	}

	return containersToStart, companionStart, companionID, nil
}

func createDefaultNetwork() (id string, err error) {

	client, err := DockerSDK.GetNewClient()

	if err != nil {
		return "", nil
	}

	id, err = client.CreateNetwork()
	return id, err
}

func StartContainersWithDefaultNetwork(containers map[string]string) error {

	fmt.Printf(constants.Blue + "\n🐳 Starting Required Containers for Companion\n\n" + constants.White)

	defaultNetworkId, err := createDefaultNetwork()

	if err != nil {
		return err
	}

	client, err := DockerSDK.GetNewClient()

	if err != nil {
		return err
	}
	if _, ok := containers[constants.MongoDBImage]; ok {
		mongodbID, err := DefaultContainers.LaunchMongoDbContainer(*client, defaultNetworkId)

		if err != nil {
			fmt.Println(Utils.Cross() + "Error Starting Mongodb Container for Rocket.Chat, Aborting Operation")
			return err
		}

		fmt.Println(Utils.Tick() + "Started Mongodb Container for Rocket.Chat with ID " + mongodbID)
	}

	if _, ok := containers[constants.RocketChatImage]; ok {
		rocketChatID, err := DefaultContainers.LaunchRocketChatContainer(*client, defaultNetworkId)

		if err != nil {
			fmt.Println(Utils.Cross() + "Error Starting Rocket.Chat Container for Rocket.Chat, Aborting Operation")
			return err
		}

		fmt.Println(Utils.Tick() + "Started Rocket.Chat Container with ID " + rocketChatID)
	}

	time.Sleep(10 * time.Second)

	return nil
}

func CreateAdminUser() error {

	fmt.Printf(constants.Blue + "\n🐳 Creating Admin User for Rocket.Chat Instance\n\n" + constants.White)

	initiateadmin.Initiate()

	return nil
}

func StartCompanionContainer(appDir string) error {

	fmt.Printf(constants.Blue + "\n🐳 Starting Required Containers for Companion\n\n" + constants.White)

	defaultNetworkId, err := createDefaultNetwork()

	if err != nil {
		return err
	}

	client, err := DockerSDK.GetNewClient()

	if err != nil {
		return err
	}

	_, err = DefaultContainers.LaunchCompanionContainer(*client, defaultNetworkId, appDir)

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
