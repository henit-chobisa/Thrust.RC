package DockerSDK

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	models "thrust/Packages/Models"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
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

func (d *Docker) CreateContainer(containerDesc models.Container, showLogs bool) (containerID string, err error) {

	cont, err := d.Client.ContainerCreate(context.TODO(), &container.Config{
		Image:        containerDesc.Image,
		Cmd:          containerDesc.Commands,
		AttachStdout: containerDesc.Stdout,
		ExposedPorts: containerDesc.ExposedPort,
		Env:          containerDesc.Env,
		Volumes:      containerDesc.Volumes,
	}, &container.HostConfig{
		// AutoRemove:   true,
		Binds:        containerDesc.Binds,
		Mounts:       containerDesc.Mount,
		PortBindings: containerDesc.PortBindings,
		RestartPolicy: container.RestartPolicy{
			Name:              "on-failure",
			MaximumRetryCount: 2,
		},
		Links:       containerDesc.Links,
		NetworkMode: "bridge",
	}, &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"RCAPPSDEFAULT": {
				Aliases: containerDesc.Aliases,
			},
		},
	}, nil, containerDesc.ContainerName)

	if err != nil {
		return "", err
	}

	err = d.Client.NetworkConnect(context.TODO(), containerDesc.NetworkID, cont.ID, nil)

	if err != nil {
		fmt.Println("Network Connection Error")
		return "", err
	}

	if err := d.Client.ContainerStart(context.TODO(), cont.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println("Container Start Error")
		return "", err
	}

	if showLogs {
		out, err := d.Client.ContainerLogs(context.Background(), cont.ID, types.ContainerLogsOptions{ShowStdout: true, Follow: true})
		if err != nil {
			return "", err
		}
		io.Copy(os.Stdout, out)
	}

	return containerID, err
}
