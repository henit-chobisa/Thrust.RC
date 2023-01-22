package DefaultContainers

import (
	constants "RCTestSetup/Packages/Constants"
	"RCTestSetup/Packages/DockerSDK"
	"RCTestSetup/Utils"
)

func LaunchCompanionContainer(sdk DockerSDK.Docker, networkID string) (containerID string, err error) {
	containerID, err = sdk.CreateContainer(networkID, "companion_"+Utils.RandomString(5), constants.CompanionImage, nil, nil, nil, nil, []string{
		"/workspace/RC-Test-Environment-Companion/github:/app",
	}, []string{"Companion"}, nil, nil, true, []string{"watch", "--url", "http://rocketchat:3000/", "--username", "user0", "--password", "123456"}, true)

	return
}
