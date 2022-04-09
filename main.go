package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	retro "github.com/necatican/agile-tools/pkg/apis/retro/v1alpha1"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	s := grpc.NewServer()

	retro.RegisterServers(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
