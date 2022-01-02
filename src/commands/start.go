package commands

import (
	"context"
	"devEnv/src/container"
	"errors"
	"github.com/docker/docker/api/types"
)

func StartCommand(p Params) error {

	containerName := p.GetValue("-name")
	if containerName == "" {
		return errors.New("invalid container name")
	}

	cli, err := container.GetConnection()
	if err != nil {
		return err
	}
	ctx := context.Background()
	options := types.ContainerStartOptions{}
	err = cli.ContainerStart(ctx, containerName, options)
	if err != nil {
		return err
	}

	return nil
}
