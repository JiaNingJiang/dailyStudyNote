package main

import (
	"context"
	"fmt"
	"github.com/go-grpc-example/proto/search"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"google.golang.org/grpc"
)

const PORT = "9001"

func main() {

	var authInterceptor grpc.UnaryServerInterceptor // 定义一个拦截器用来做token验证
	authInterceptor = func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		err = Auth(ctx) // 拦截普通方法请求，验证 Token
		if err != nil { // 验证失败，直接退出
			return
		}
		return handler(ctx, req) // 如果验证成功，则继续处理请求
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	serviceObject := new(search.SearchService)
	search.RegisterSearchServiceServer(server, serviceObject)

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}

func Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var user string
	var password string

	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}

	if user != "admin" || password != "admin" {
		return status.Errorf(codes.Unauthenticated, "token不合法")
	}
	return nil
}
