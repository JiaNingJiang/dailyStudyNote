package server

import (
	"context"
	"crypto/tls"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-gateway-example/pkg/utils"
	"github.com/grpc-gateway-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"
)

var (
	ServerPort     string
	ServerCertName string
	ServerPemPath  string
	ServerKeyPath  string
	EndPoint       string
)

func Serve() (err error) {
	EndPoint = ":" + ServerPort
	conn, err := net.Listen("tcp", EndPoint)
	if err != nil {
		log.Printf("TCP Listen err:%v\n", err)
	}

	tlsConfig := utils.GetTLSConfig(ServerPemPath, ServerKeyPath)
	srv := createInternalServer(conn, tlsConfig)

	log.Printf("gRPC and https listen on: %s\n", ServerPort)

	if err = srv.Serve(tls.NewListener(conn, tlsConfig)); err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}

	return err
}

func createInternalServer(conn net.Listener, tlsConfig *tls.Config) *http.Server {
	var opts []grpc.ServerOption

	creds, err := credentials.NewServerTLSFromFile(ServerPemPath, ServerKeyPath) // 根据服务端证书和私钥构造服务端的TLS证书凭证
	if err != nil {
		log.Printf("Failed to create server TLS credentials %v", err)
	}

	opts = append(opts, grpc.Creds(creds))
	grpcServer := grpc.NewServer(opts...) // 创建一个需要TLS验证的grpc服务器

	proto.RegisterHelloWorldServer(grpcServer, NewHelloService()) // 为grpc服务器注册服务

	ctx := context.Background()
	dcreds, err := credentials.NewClientTLSFromFile(ServerPemPath, ServerCertName) // 根据服务端证书和服务域名构造客户端的TLS证书凭证
	if err != nil {
		log.Printf("Failed to create client TLS credentials %v", err)
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	gwmux := runtime.NewServeMux()

	// register grpc-gateway pb
	if err := proto.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, EndPoint, dopts); err != nil {
		log.Printf("Failed to register gw server: %v\n", err)
	}

	// http服务
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	return &http.Server{
		Addr:      EndPoint,
		Handler:   utils.GrpcHandlerFunc(grpcServer, mux),
		TLSConfig: tlsConfig,
	}
}

func ServeWithoutTLS() (err error) {
	EndPoint = ":" + ServerPort
	server := grpc.NewServer() // 没有添加TLS证书认证

	proto.RegisterHelloWorldServer(server, NewHelloService())

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}

	err = proto.RegisterHelloWorldHandlerFromEndpoint(context.Background(), gwmux, EndPoint, dopts)
	if err != nil {
		log.Printf("Failed to register gw server: %v\n", err)
	}

	mux.Handle("/", gwmux)
	http.ListenAndServe(EndPoint, utils.GrpcHandlerFunc(server, mux))

	return
}
