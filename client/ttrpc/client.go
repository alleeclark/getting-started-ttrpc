package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"ttrpc-demo/pb/hello"

	"github.com/containerd/ttrpc"
	"github.com/sirupsen/logrus"
)

const port = ":9001"

func main() {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to dial: %v \n", err)
		os.Exit(1)
	}
	client := hello.NewHelloServiceClient(ttrpc.NewClient(conn))
	serverResponse, err := client.HelloWorld(context.Background(), &hello.HelloRequest{
		Msg: "Hello Server",
	})
	if err != nil {
		logrus.WithError(err).Error("failed connecting to server")
		os.Exit(1)
	}
	logrus.WithField("resp", serverResponse.Response).Debug("recieved request from server")
}
