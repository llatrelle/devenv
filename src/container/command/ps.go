package command

import (
	"context"
	"devEnv/src/api/model"
	"devEnv/src/container"
	"github.com/docker/docker/api/types"
	"log"
)

//PS List docker containers
func PS(p Params) ([]model.Environment, error) {
	var envs []model.Environment
	cli, err := container.GetConnection()
	if err != nil {
		log.Printf("connection error...")
	}
	ctx := context.Background()
	options := types.ContainerListOptions{All: p.HaveArg("-all")}
	containers, err := cli.ContainerList(ctx, options)

	if err != nil {
		return envs, err
	}
	log.Printf("Containers: %v", len(containers))

	for _, c := range containers {
		e := model.Environment{
			Name:   c.Names[0],
			Image:  c.Image,
			Status: c.Status,
		}
		envs = append(envs, e)
		//log.Printf("ID: %v State: %v Status: %v, Name: %v", c.ID, c.State, c.Status, c.Names[0])
	}

	return envs, nil
}
