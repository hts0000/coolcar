package mongotesting

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type URIConfig struct {
	IP   string
	Port string
}

const (
	image         = "mongo:4.4"
	containerPort = "27017/tcp"
)

func RunWithMongoInDocker(m *testing.M, config URIConfig, mongoURI *string) int {
	c, err := client.NewClientWithOpts(client.WithHost("tcp://" + config.IP + ":" + config.Port))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	resp, err := c.ContainerCreate(ctx, &container.Config{
		ExposedPorts: nat.PortSet{
			"27017/tcp": {},
		},
		Image: image,
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "0", // 随机挑选一个可用端口
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}
	containerID := resp.ID
	defer func() {
		fmt.Println("killing mongodb container...")
		err := c.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true})
		if err != nil {
			fmt.Println("kill mongodb container failed")
			panic(err)
		}
		fmt.Println("kill mongodb container successed")
	}()
	fmt.Println("starting mongodb container...")
	err = c.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("start mongodb container successed")
	inspRes, err := c.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}
	hostPort := inspRes.NetworkSettings.Ports["27017/tcp"][0]
	*mongoURI = fmt.Sprintf("mongodb://%s:%s", config.IP, hostPort.HostPort)
	// time.Sleep(time.Second * 10)

	return m.Run()
}
