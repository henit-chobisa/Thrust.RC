package DefaultContainers

import (
	constants "RCTestSetup/Packages/Constants"
	"RCTestSetup/Packages/DockerSDK"
	"RCTestSetup/Utils"

	"github.com/docker/go-connections/nat"
)

func LaunchRocketChatContainer(sdk DockerSDK.Docker, networkID string) (string, error) {
	containerID, err := sdk.CreateContainer(networkID, "rocketchat_"+Utils.RandomString(5), constants.RocketChatImage, nat.PortSet{
		"3000/tcp": {},
	}, nat.PortMap{
		"3000/tcp": []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: "3000",
			},
		},
	}, []string{
		"MONGO_URL=mongodb://mongo:27017/rocketchat?replicaSet=rs0",
		"MONGO_OPLOG_URL=mongodb://mongo:27017/local?replicaSet=rs0",
		"ROOT_URL=http://localhost:3000",
		"PORT=3000",
		"DEPLOY_METHOD=docker",
		"OVERWRITE_SETTING_Apps_Framework_Development_Mode=true",
		"OVERWRITE_SETTING_Show_Setup_Wizard=Completed",
	}, nil, nil, []string{
		"rocketchat",
	}, nil, nil, false, nil, false)

	return containerID, err
}
