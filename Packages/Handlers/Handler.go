package Handlers

import (
	constants "RCTestSetup/Packages/Constants"
	"RCTestSetup/Packages/DockerSDK"
	"RCTestSetup/Utils"
	"RCTestSetup/tui/components/Page"
	"fmt"
	"sync"

	"github.com/docker/docker/api/types/filters"
)

func HandleDependencyCheck() error {
	// confirming initial configuration dependency
	fmt.Printf(constants.Blue + "\nChecking Initial dependencies required for running the companion\n\n")
	dependencyModel := Page.NewDependencyModel()
	fmt.Println(dependencyModel.View())

	if dependencyModel.Err != nil {
		return dependencyModel.Err
	}
	return nil
}

func HandlePullingImages() (map[string]string, error) {
	// first let's check for the images that are currently present in the system

	fmt.Printf(constants.Blue + "Verifying Required Docker Images for Companion\n\n" + constants.White)

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
	fmt.Printf(constants.Blue + "\nPulling Required Images\n" + constants.White)

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

	return nil
}
