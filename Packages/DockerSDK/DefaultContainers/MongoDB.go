package DefaultContainers

import (
	constants "thrust/Packages/Constants"
	"thrust/Packages/DockerSDK"
	models "thrust/Packages/Models"
	"thrust/Utils"

	"github.com/docker/go-connections/nat"
)

func LaunchMongoDbContainer(sdk DockerSDK.Docker, networkID string) (string, error) {

	mongoDBContaier := models.Container{
		NetworkID:     networkID,
		ContainerName: "mongodb_" + Utils.RandomString(5),
		Image:         constants.MongoDBImage,
		ExposedPort: nat.PortSet{
			"27017/tcp": {},
		},
		PortBindings: nil,
		Env: []string{
			"MONGODB_REPLICA_SET_MODE=primary",
			"MONGODB_REPLICA_SET_NAME=rs0",
			"MONGODB_PORT_NUMBER=27017",
			"MONGODB_INITIAL_PRIMARY_HOST=mongodb",
			"MONGODB_INITIAL_PRIMARY_PORT_NUMBER=27017",
			"MONGODB_ADVERTISED_HOSTNAME=mongodb",
			"MONGODB_ENABLE_JOURNAL=true",
			"ALLOW_EMPTY_PASSWORD=yes",
		},
		Volumes: map[string]struct{}{
			"/bitnami/mongodb": {},
		},
		Binds: nil,
		Aliases: []string{
			"mongodb",
			"mongo",
		},
		Links:      nil,
		Mount:      nil,
		Commands:   nil,
		Stdout:     false,
		AutoRemove: false,
	}

	constainerID, err := sdk.CreateContainer(mongoDBContaier, false)

	return constainerID, err
}
