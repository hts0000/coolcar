package main

import (
	"context"
	"coolcar/car/mq/amqpclt"
	coolenvpb "coolcar/shared/coolenv"
	"coolcar/shared/server"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("hts0000.top:18001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	ac := coolenvpb.NewAIServiceClient(conn)
	c := context.Background()

	// 计算距离
	res, err := ac.MeasureDistance(c, &coolenvpb.MeasureDistanceRequest{
		From: &coolenvpb.Location{
			Latitude:  23.15792,
			Longitude: 113.27324,
		},
		To: &coolenvpb.Location{
			Latitude:  23.15892,
			Longitude: 113.27424,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)

	// 驾照信息识别
	idRes, err := ac.LicIdentity(c, &coolenvpb.IdentityRequest{
		Photo: []byte{1, 2, 3, 4, 5},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", idRes)

	// 模拟位置更新，服务端会不断的往服务器上的RabbitMQ上的exchange发送位置信息
	_, err = ac.SimulateCarPos(c, &coolenvpb.SimulateCarPosRequest{
		CarId: "car123",
		InitialPos: &coolenvpb.Location{
			Latitude:  30,
			Longitude: 120,
		},
		Type: coolenvpb.PosType_RANDOM,
	})
	if err != nil {
		panic(err)
	}

	logger, err := server.NewZapLogger()
	if err != nil {
		panic(err)
	}

	amqpConn, err := amqp.Dial("amqp://guest:guest@hts0000.top:5672/")
	if err != nil {
		panic(err)
	}
	defer amqpConn.Close()

	sub, err := amqpclt.NewSubscriber(amqpConn, "pos_sim", logger)
	if err != nil {
		panic(err)
	}

	ch, cleanUp, err := sub.SubscribeRaw(c)
	defer cleanUp()
	if err != nil {
		panic(err)
	}

	tm := time.After(10 * time.Second)
	for {
		shouldStop := false
		select {
		case msg := <-ch:
			fmt.Printf("%s\n", msg.Body)
		case <-tm:
			shouldStop = true
		}
		if shouldStop {
			break
		}
	}

	_, err = ac.EndSimulateCarPos(c, &coolenvpb.EndSimulateCarPosRequest{
		CarId: "car123",
	})
	if err != nil {
		panic(err)
	}
}
