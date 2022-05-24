package service

import (
	"context"
	"demo_grpc_server/protos/everphone"
	"fmt"
	"google.golang.org/grpc"
)

type server struct {
	everphone.EverphoneServer
}

func (s *server) RandomText(ctx context.Context, in *everphone.EverphoneRandomTextInput) (*everphone.EverphoneRandomTextOutput, error) {
	return &everphone.EverphoneRandomTextOutput{
		Text: fmt.Sprintf("Random text: %s", in.Text),
	}, nil
}

func New(grpcServer *grpc.Server) {
	everphone.RegisterEverphoneServer(grpcServer, &server{})
}
