package grpcmethods

import (
	"context"

	"github.com/jamolh/chatik-sharedpb/service/chatikpb"
	"google.golang.org/grpc"
)

type ChatikServer struct{}

func RegisterService(server *grpc.Server) {
	chatikpb.RegisterChatikServiceServer(server, &ChatikServer{})
}

func (*ChatikServer) Ping(ctx context.Context, request *chatikpb.PingRequest) (*chatikpb.PingResponse, error) {
	return &chatikpb.PingResponse{
		Message: "i'm ok!",
	}, nil
}
