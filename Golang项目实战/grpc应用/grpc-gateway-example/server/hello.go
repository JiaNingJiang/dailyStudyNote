package server

import (
	"context"
	"github.com/grpc-gateway-example/proto"
)

type helloService struct{}

func NewHelloService() *helloService {
	return &helloService{}
}

func (h helloService) SayHelloWorld(ctx context.Context, r *proto.HelloWorldRequest) (*proto.HelloWorldResponse, error) {
	return &proto.HelloWorldResponse{
		Message: "This is grpc SayHelloWorld",
	}, nil
}

func (h helloService) MustEmbedUnimplementedHelloWorldServer() {

}
