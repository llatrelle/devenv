package command

import (
	"context"
	"devEnv/src/container"
	"errors"
	"time"
)

func Stop(p Params) error {
	var err error
	containerName := p.GetValue("-name")
	timeout, _ := time.ParseDuration("30s")
	if p.HaveArg("-t") {
		timeout, err = time.ParseDuration(p.GetValue("-t"))
	}
	if err != nil {
		return errors.New("invalid time out, format is 10h20m30s or 30s")
	}
	if containerName == "" {
		return errors.New("invalid container name")
	}

	cli, err := container.GetConnection()
	if err != nil {
		return err
	}
	ctx := context.Background()
	err = cli.ContainerStop(ctx, containerName, &timeout)
	if err != nil {
		return err
	}

	return nil
}
