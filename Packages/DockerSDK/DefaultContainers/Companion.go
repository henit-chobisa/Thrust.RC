package DefaultContainers

import (
	constants "RCTestSetup/Packages/Constants"
	"RCTestSetup/Packages/DockerSDK"
	"RCTestSetup/Utils"
)

func LaunchCompanionContainer(sdk DockerSDK.Docker, networkID string) (containerID string, err error) {
	containerID, err = sdk.CreateContainer(networkID, "companion_"+Utils.RandomString(5), constants.CompanionImage, nil, nil, nil, nil, []string{
		"./:/app",
	}, []string{"Companion"}, nil, nil)

	return
}
