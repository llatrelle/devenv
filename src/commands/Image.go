package commands

import (
	"context"
	"devEnv/src/container"
	"errors"
	"github.com/docker/docker/api/types"
	"io"
	"os"
)

func PullImageCommand(p Params) error {

	ctx := context.Background()
	image := p.GetValue("-i")
	if image == "" {
		return errors.New("invalid image name")
	}
	cli, err := container.GetConnection()
	if err != nil {
		return err
	}

	reader, err := cli.ImagePull(ctx, p.GetValue("-i"), types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader)
	return nil
}

func ImageCommand(p Params) error {
	if p.HaveArg("-pull") {
		return PullImageCommand(p)
	}

	return errors.New("invalid command")
}
