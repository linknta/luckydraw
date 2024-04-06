package luckydraw

import (
	pb "github.com/linknta/luckydraw/api/v1/luckydraw"
	"github.com/linknta/luckydraw/internal/service/machine"
)

type luckydrawServer struct {
	machine machine.Machine

	pb.UnimplementedLuckydrawServer
}

func NewLuckydrawServer(machine machine.Machine) pb.LuckydrawServer {
	return &luckydrawServer{
		machine: machine,
	}
}
