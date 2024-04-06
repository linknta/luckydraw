package luckydraw

import (
	"context"

	pb "github.com/linknta/luckydraw/api/v1/luckydraw"
)

func (s *luckydrawServer) Roll(ctx context.Context, request *pb.RollRquest) (*pb.RollReply, error) {
	// TODO: validate

	result, err := s.machine.Play()
	if err != nil {
		return nil, err
	}

	// TODO: handle state

	return &pb.RollReply{
		Result: result,
	}, nil
}
