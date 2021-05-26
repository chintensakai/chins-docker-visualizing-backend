package sdk

import (
	"github.com/docker/docker/api/types"
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

	// ctx := context.Background()

	// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation(),
	// 	client.WithHost("tcp://192.168.1.248:2375"))

	// 本地
	// cli, err := client.NewClientWithOpts(client.FromEnv)
	// if err != nil {
	// 	panic(err)
	// }
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}
	result := make([]Container, 0, len(containers))
	for _, container := range containers {
		tmp := Container{
			ID:      container.ID[:16],
			Names:   container.Names[0],
			Image:   container.Image,
			ImageID: container.ImageID[8:24],
			Created: container.Created,
			State:   container.State,
		}
		result = append(result, tmp)
	}
	return result
}

func StartContainer(containerId string) {
	if err := cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}

func StopContainer(containerId string) {
	if err := cli.ContainerStop(ctx, containerId, nil); err != nil {
		panic(err)
	}
}
