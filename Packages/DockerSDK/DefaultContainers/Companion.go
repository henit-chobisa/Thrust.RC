package DefaultContainers

import (
	"fmt"
	constants "thrust/Packages/Constants"
	"thrust/Packages/DockerSDK"
	models "thrust/Packages/Models"

	"github.com/spf13/viper"
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
			"username=" + viper.GetString("admin.username"),
			"password=" + viper.GetString("admin.password"),
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
