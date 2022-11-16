package main

import (
	"context"
	carpb "coolcar/car/api/gen/v1"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	cs := carpb.NewCarServiceClient(conn)
	c := context.Background()
	for i := 0; i < 5; i++ {
		ce, err := cs.CreateCar(c, &carpb.CreateCarRequest{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("create car: %s\n", ce.Id)
	}
}
