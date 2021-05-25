package sdk

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Image struct {
	Id         string `json:"Id"`
	Created    int64
	Repository string
	Tags       string
	Size       int64
}

func GetImages() []Image {
	ctx := context.Background()

	// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation(),
	// 	client.WithHost("tcp://192.168.1.248:2375"))

	// 本地
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})

	result := make([]Image, 0, len(images))
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		tmp := Image{
			Id:         image.ID,
			Created:    image.Created,
			Repository: strings.Split(image.RepoTags[0], ":")[0],
			Tags:       strings.Split(image.RepoTags[0], ":")[1],
			Size:       image.Size,
		}
		result = append(result, tmp)
	}
	return result
}
