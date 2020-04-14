package main

import (
	"context"
	"net"
	"os"
	"ttrpc-demo/pb/hello"

	"github.com/containerd/ttrpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const port = ":9000"

func main() {

	srv, err := ttrpc.NewServer()
	defer srv.Close()
	if err != nil {
		logrus.WithError(err).Error("failed starting a new server")
		os.Exit(1)
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.WithError(err).WithField("port", port).Error("failed to listen on port")
		os.Exit(1)
	}
	hello.RegisterHelloServiceService(srv, &helloService{})
	if err := srv.Serve(context.Background(), lis); err != nil {
		logrus.WithError(err).WithField("port", port).Error("failed serving connection")
		os.Exit(1)

	}
}

type helloService struct{}

func (s helloService) HelloWorld(ctx context.Context, r *hello.HelloRequest) (*hello.HelloResponse, error) {
	if r.Msg == "" {
		return nil, status.Error(codes.InvalidArgument, "ErrNoInputMsgGiven")
	}
	return &hello.HelloResponse{Response: "Hi How are you"}, nil
}
