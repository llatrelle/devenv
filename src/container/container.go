package container

import (
	"errors"
	"github.com/docker/docker/client"
	"sync"
)

var (
	conn  *client.Client
	mutex sync.Mutex
)

//GetConnection return a docker connection
func GetConnection() (*client.Client, error) {
	mutex.Lock()
	if conn == nil {
		c, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			return nil, errors.New("error connecting with docker")
		}
		conn = c
	}
	mutex.Unlock()
	return conn, nil

}
