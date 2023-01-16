package DockerSDK

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
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

func (d *Docker) FindImages(filters filters.Args) (*[]types.ImageSummary, error) {
	images, err := d.client.ImageList(context.TODO(), types.ImageListOptions{
		All:     true,
		Filters: filters,
	})

	if err != nil {
		return nil, err
	}
	return &images, err

}

func (d *Docker) FindContainers(filters filters.Args) (*[]types.Container, error) {
	containers, err := d.client.ContainerList(context.TODO(), types.ContainerListOptions{
		Size:    true,
		All:     true,
		Since:   "Container",
		Filters: filters,
	})

	if err != nil {
		return nil, err
	}
	return &containers, err
}
