package DefaultContainers

import (
	"fmt"
	constants "thrust/Packages/Constants"
	"thrust/Packages/DockerSDK"
	"thrust/Utils"
)

func LaunchCompanionContainer(sdk DockerSDK.Docker, networkID string, path string, name string) (containerID string, err error) {
	containerID, err = sdk.CreateContainer(networkID, name+Utils.RandomString(5), constants.CompanionImage, nil, nil, []string{
		"url=http://rocketchat:3000",
		"username=user0",
		"password=123456",
	}, nil, []string{
		fmt.Sprintf("%v:/app", path),
	}, []string{"Companion"}, nil, nil, true, nil, true)

	return
}
