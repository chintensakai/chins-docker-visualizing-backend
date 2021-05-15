package sdk

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Container struct {
	ID      string `json:"Id"`
	Names   string
	Image   string
	ImageID string
	Created int64
	State   string
}

func GetContainers() []Container {

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation(),
		client.WithHost("tcp://192.168.1.248:2375"))

	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}
	result := make([]Container, 0, len(containers))
	for _, container := range containers {
		tmp := Container{
			ID:      container.ID,
			Names:   container.Names[0],
			Image:   container.Image,
			ImageID: container.ImageID,
			Created: container.Created,
			State:   container.State,
		}
		result = append(result, tmp)
	}
	return result
}
