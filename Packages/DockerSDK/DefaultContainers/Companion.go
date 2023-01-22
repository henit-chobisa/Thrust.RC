package DefaultContainers

import (
	constants "AppsCompanion/Packages/Constants"
	"AppsCompanion/Packages/DockerSDK"
	"AppsCompanion/Utils"
	"fmt"
)

func LaunchCompanionContainer(sdk DockerSDK.Docker, networkID string, path string) (containerID string, err error) {
	containerID, err = sdk.CreateContainer(networkID, "companion_"+Utils.RandomString(5), constants.CompanionImage, nil, nil, []string{
		"url=http://rocketchat:3000",
		"username=user0",
		"password=123456",
	}, nil, []string{
		fmt.Sprintf("%v:/app", path),
	}, []string{"Companion"}, nil, nil, true, nil, true)

	return
}
