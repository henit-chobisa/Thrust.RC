package models

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
)

type Container struct {
	NetworkID     string
	ContainerName string
	Image         string
	ExposedPort   nat.PortSet
	PortBindings  nat.PortMap
	Env           []string
	Volumes       map[string]struct{}
	Binds         []string
	Aliases       []string
	Links         []string
	Mount         []mount.Mount
	Commands      []string
	Stdout        bool
	AutoRemove    bool
}
