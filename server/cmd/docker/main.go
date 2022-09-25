package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	c, err := client.NewClientWithOpts(client.WithHost("tcp://182.61.47.223:9876"))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	resp, err := c.ContainerCreate(ctx, &container.Config{
		ExposedPorts: nat.PortSet{
			"27017/tcp": {},
		},
		Image: "mongo:4.4",
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"27017/tcp": []nat.PortBinding{
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
	fmt.Println("starting mongodb container...")
	err = c.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Println("killing mongodb container...")
		err := c.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{Force: true})
		if err != nil {
			fmt.Println("kill mongodb container failed")
		} else {
			fmt.Println("kill mongodb container successed")
		}
	}()
	fmt.Println("start mongodb container successed")
	inspRes, err := c.ContainerInspect(ctx, resp.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listening: %v:%v\n", inspRes.NetworkSettings.Ports["27017/tcp"][0].HostIP, inspRes.NetworkSettings.Ports["27017/tcp"][0].HostPort)
	time.Sleep(time.Second * 10)
}
