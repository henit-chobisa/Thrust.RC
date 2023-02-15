package DefaultContainers

import (
	"fmt"
	constants "thrust/Packages/Constants"
	"thrust/Packages/DockerSDK"
	models "thrust/Packages/Models"
)

func LaunchCompanionContainer(sdk DockerSDK.Docker, networkID string, path string, name string) (containerID string, err error) {

	companionContainer := models.Container{
		NetworkID:     networkID,
		ContainerName: name,
		Image:         constants.CompanionImage,
		ExposedPort:   nil,
		PortBindings:  nil,
		Env: []string{
			"url=http://rocketchat:3000",
			"username=user0",
			"password=123456",
		},
		Volumes: nil,
		Binds: []string{
			fmt.Sprintf("%v:/app", path),
		},
		Aliases:    []string{"Companion"},
		Links:      nil,
		Mount:      nil,
		Commands:   nil,
		Stdout:     true,
		AutoRemove: true,
	}

	containerID, err = sdk.CreateContainer(companionContainer, true)

	return
}
