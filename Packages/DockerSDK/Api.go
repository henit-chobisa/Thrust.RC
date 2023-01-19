package DockerSDK

import (
	"context"
	"io"
	"os"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type Docker struct {
	Client  *client.Client
	images  []string
	network types.NetworkResource
}

func GetVersionInfo() (*types.Version, error) {
	c, err := GetNewClient()
	if err != nil {
		return nil, err
	}
	version, err := c.Client.ServerVersion(context.Background())
	if err != nil {
		return nil, err
	}
	c.Client.Close()
	return &version, nil
}

func GetNewClient() (*Docker, error) {
	c, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	docker := Docker{
		Client: c,
	}
	return &docker, nil
}

func (d *Docker) FindImages(filters filters.Args) (*[]types.ImageSummary, error) {
	images, err := d.Client.ImageList(context.TODO(), types.ImageListOptions{
		All:     true,
		Filters: filters,
	})

	if err != nil {
		return nil, err
	}
	return &images, err

}

func (d *Docker) FindContainers(filters filters.Args) (*[]types.Container, error) {
	containers, err := d.Client.ContainerList(context.TODO(), types.ContainerListOptions{
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

func (d *Docker) PullImage(image string, wg *sync.WaitGroup) error {
	defer wg.Done()
	out, err := d.Client.ImagePull(context.TODO(), image, types.ImagePullOptions{})

	if err != nil {
		wg.Done()
		return err
	}

	io.Copy(os.Stdout, out)
	out.Close()
	return nil
}
