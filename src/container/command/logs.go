package command

import (
	"context"
	"devEnv/src/container"
	"errors"
	"github.com/docker/docker/api/types"
	"io"
	"io/ioutil"
	"os"
)

//Logs return in os.Stdout logs from container
func Logs(p Params) error {

	containerName := p.GetValue("-name")
	if containerName == "" {
		return errors.New("invalid container name")
	}

	cli, err := container.GetConnection()
	if err != nil {
		return err
	}
	ctx := context.Background()
	options := types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: p.HaveArg("-f")}
	reader, err := cli.ContainerLogs(ctx, containerName, options)

	//defer reader.Close()
	if err != nil {
		return err
	}
	ioutil.ReadAll(reader)
	io.Copy(os.Stdout, reader)

	return nil
}
