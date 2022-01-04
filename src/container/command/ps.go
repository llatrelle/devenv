package command

import (
	"context"
	"devEnv/src/container"
	"github.com/docker/docker/api/types"
	"log"
)

//PS List docker containers
func PS(p Params) error {
	cli, err := container.GetConnection()
	if err != nil {
		log.Printf("connection error...")
	}
	ctx := context.Background()
	options := types.ContainerListOptions{All: p.HaveArg("-all")}
	containers, err := cli.ContainerList(ctx, options)
	if err != nil {
		return err
	}
	log.Printf("Containers: %v", len(containers))
	for _, c := range containers {
		log.Printf("ID: %v State: %v Status: %v, Name: %v", c.ID, c.State, c.Status, c.Names[0])
	}
	return nil
}
