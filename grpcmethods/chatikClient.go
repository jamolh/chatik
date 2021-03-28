package grpcmethods

import (
	"github.com/jamolh/chatik-sharedpb/service/chatikpb"
	"google.golang.org/grpc"
)

var (
	conn *grpc.ClientConn
)

func Connect(addr string) (chatikpb.ChatikServiceClient, error) {
	var err error
	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	return chatikpb.NewChatikServiceClient(conn), err
}

func Disconnect() error {
	return conn.Close()
}
