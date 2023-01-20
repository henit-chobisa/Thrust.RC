package DockerSDK

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type Docker struct {
	Client *client.Client
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

func (d *Docker) NetworkExist(filters filters.Args, name string) (string, bool, error) {
	networks, err := d.Client.NetworkList(context.Background(), types.NetworkListOptions{
		Filters: filters,
	})

	if err != nil {
		return "", false, err
	}

	for _, network := range networks {
		if name == network.Name {
			return network.ID, true, nil
		}
	}

	return "", false, nil
}

func (d *Docker) CreateNetwork() (string, error) {

	filter := filters.NewArgs()
	filter.Add("name", "RCAPPSDEFAULT")

	networkID, networkExist, err := d.NetworkExist(filter, "RCAPPSDEFAULT")

	if err != nil {
		return "", err
	}

	if networkExist {
		return networkID, nil
	}

	response, err := d.Client.NetworkCreate(context.TODO(), "RCAPPSDEFAULT", types.NetworkCreate{
		CheckDuplicate: true,
		Driver:         "bridge",
		Options: map[string]string{
			"com.docker.network.bridge.enable_icc": "true",
			"com.docker.network.driver.mtu":        "1440",
		},
	})
	if err != nil {
		return "", err
	}
	return response.ID, nil
}

func (d *Docker) CreateContainer(networkId string, name string, image string, portExpose nat.PortSet, portBindings nat.PortMap, env []string, volumes map[string]struct{}, binds []string, aliases []string, links []string, mounts []mount.Mount) (containerID string, err error) {

	container, err := d.Client.ContainerCreate(context.TODO(), &container.Config{
		Image:        image,
		AttachStdout: false,
		ExposedPorts: portExpose,
		Env:          env,
		Volumes:      volumes,
	}, &container.HostConfig{
		// AutoRemove:   true,
		Binds:        binds,
		Mounts:       mounts,
		PortBindings: portBindings,
		RestartPolicy: container.RestartPolicy{
			Name:              "on-failure",
			MaximumRetryCount: 5,
		},
		Links:       links,
		NetworkMode: "bridge",
	}, &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"RCAPPSDEFAULT": {
				Aliases: aliases,
			},
		},
	}, nil, name)

	if err != nil {
		return "", err
	}

	err = d.Client.NetworkConnect(context.TODO(), networkId, container.ID, nil)

	if err != nil {
		fmt.Println("Network Connection Error")
		return "", err
	}

	if err := d.Client.ContainerStart(context.TODO(), container.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println("Container Start Error")
		return "", err
	}

	return containerID, err
}
