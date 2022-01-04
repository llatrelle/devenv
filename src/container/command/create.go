package command

import (
	"context"
	"devEnv/src/container"
	dc "github.com/docker/docker/api/types/container"
	"log"
)

func Create(p Params) error {
	containerName := p.GetValue("-name")
	cli, err := container.GetConnection()
	if err != nil {
		return err
	}
	ctx := context.Background()
	options := dc.Config{
		Image: p.GetValue("-i"),
		Cmd:   []string{"echo", "hello world"},
	}

	resp, err := cli.ContainerCreate(ctx, &options, nil, nil, nil, containerName)
	if err != nil {
		return nil
	}
	log.Printf("ContainerID:/n %v", resp.ID)
	log.Printf("Warrnings:/n %v", resp.Warnings)
	return nil
}
