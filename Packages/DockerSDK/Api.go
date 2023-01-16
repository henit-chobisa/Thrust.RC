package DockerSDK

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Docker struct {
	client  *client.Client
	images  []string
	network types.NetworkResource
}

func GetVersionInfo() (*types.Version, error) {
	c, err := GetNewClient()
	if err != nil {
		return nil, err
	}
	version, err := c.client.ServerVersion(context.Background())
	if err != nil {
		return nil, err
	}
	c.client.Close()
	return &version, nil
}

func GetNewClient() (*Docker, error) {
	c, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	docker := Docker{
		client: c,
	}
	return &docker, nil
}
