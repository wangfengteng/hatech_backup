package client

import (
	"github.com/docker/docker/client"
)

var defaultDockerClient *client.Client

func GetDockerClient() *client.Client {
	if defaultDockerClient == nil {
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err.Error())
		}
		defaultDockerClient = cli
	}

	return defaultDockerClient
}
