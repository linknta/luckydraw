package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/linknta/luckydraw/api/v1/luckydraw"
	"github.com/linknta/luckydraw/internal/server/luckydraw"
	"github.com/linknta/luckydraw/internal/service/machine"
)

const port = 8080

func Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// register server
	var opts []grpc.ServerOption // TODO

	machine := machine.New()

	server := grpc.NewServer(opts...)
	pb.RegisterLuckydrawServer(server, luckydraw.NewLuckydrawServer(machine))

	reflection.Register(server)
	server.Serve(lis)
}
