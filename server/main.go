package main

import (
	"coolcar/trippb"
	"coolcar/tripservice"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	trippb.RegisterTripServiceServer(server, &tripservice.Service{})
	log.Fatal(server.Serve(listener))
}
