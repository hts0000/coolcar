package mongotesting

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 提供 docker 服务端的 URI 信息
type URIConfig struct {
	IP   string
	Port string
}

const (
	image         = "mongo:4.4"
	containerPort = "27017/tcp"
)

// 各个微服务的单元测试调用 RunWithMongoInDocker 时会自动设置上
// 随机创建一个新的 mongodb container 来测试
var mongoURI string

// 生产 mongo 库的ip:port
const defaultMongoURI = "mongodb://182.61.47.223:27017"

func RunWithMongoInDocker(m *testing.M, config URIConfig) int {
	// 创建 config.IP:config.Port 服务器上的docker客户端
	c, err := client.NewClientWithOpts(client.WithHost("tcp://" + config.IP + ":" + config.Port))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	// 创建 contanier
	resp, err := c.ContainerCreate(ctx, &container.Config{
		// 容器的 27017 端口需要对外开放
		ExposedPorts: nat.PortSet{
			"27017/tcp": {},
		},
		Image: image,
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{
				{
					HostIP: "0.0.0.0",
					// 对外暴露的端口在这里设置
					HostPort: "0", // 随机挑选一个可用端口
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}
	containerID := resp.ID
	// 测试结束后强制删除 container，不会影响生产库
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
	// 找到随机生成的端口
	hostPort := inspRes.NetworkSettings.Ports["27017/tcp"][0]
	// 设置上全局变量
	mongoURI = fmt.Sprintf("mongodb://%s:%s", config.IP, hostPort.HostPort)
	// time.Sleep(time.Second * 10)
	// mongodb container 开始运行，等待连接
	return m.Run()
}

// 获取随机生成的 mongodb container 客户端
func NewClient(c context.Context) (*mongo.Client, error) {
	if mongoURI == "" {
		return nil, fmt.Errorf("mongo uri not set. Please run RunWithMongoInDocker in TestMain")
	}
	fmt.Printf("use mongoURI: %q\n", mongoURI)
	return mongo.Connect(c, options.Client().ApplyURI(mongoURI))
}

// 获取生产环境的 mongodb container 客户端
func NewDefaultClient(c context.Context) (*mongo.Client, error) {
	return mongo.Connect(c, options.Client().ApplyURI(defaultMongoURI))
}

// 设置提前规划好的索引
func SetupIndexs(c context.Context, db *mongo.Database) error {
	// account 表设置索引
	_, err := db.Collection("account").Indexes().CreateOne(c, mongo.IndexModel{
		Keys: bson.D{
			// 根据 open_id 设置索引
			{
				Key: "open_id",
				// 索引从小到大
				Value: 1,
			},
		},
		// 索引不可重复
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	// trip 表设置索引
	_, err = db.Collection("trip").Indexes().CreateOne(c, mongo.IndexModel{
		Keys: bson.D{
			{
				Key:   "trip.accountid",
				Value: 1,
			},
			{
				Key:   "trip.status",
				Value: 1,
			},
		},
		Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
			"trip.status": 1,
		}),
	})
	if err != nil {
		return err
	}

	// profile 表设置索引
	_, err = db.Collection("profile").Indexes().CreateOne(c, mongo.IndexModel{
		Keys: bson.D{
			{
				Key:   "accountid",
				Value: 1,
			},
		},
		Options: options.Index().SetUnique(true),
	})
	return err
}
