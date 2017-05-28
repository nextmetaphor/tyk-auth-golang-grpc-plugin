package main

import (
	log "github.com/Sirupsen/logrus"
	coprocess "github.com/TykTechnologies/tyk-protobuf/bindings/go"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Listening...")
	s := grpc.NewServer()
	coprocess.RegisterDispatcherServer(s, &Dispatcher{})
	s.Serve(lis)
}
