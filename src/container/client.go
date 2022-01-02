package container

import "github.com/docker/docker/client"

//Container Abstraction interface for containers
type Container interface {
	GetConnection() (*client.Client, error)
}
