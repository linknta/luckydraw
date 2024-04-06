package client

import (
	"google.golang.org/grpc"

	pb "github.com/linknta/luckydraw/api/v1/luckydraw"
)

type LuckydrawClient struct {
	pb.LuckydrawClient
}

func NewLuckydrawClient(address string, opts ...grpc.DialOption) (*LuckydrawClient, error) {
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return &LuckydrawClient{
		LuckydrawClient: pb.NewLuckydrawClient(conn),
	}, nil
}
