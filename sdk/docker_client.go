package sdk

import (
	"context"

	"github.com/docker/docker/client"
)

var cli *client.Client
var ctx context.Context

func init() {
	ctx = context.Background()
	// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation(),
	// 	client.WithHost("tcp://192.168.1.248:2375"))

	// 本地
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

}
