package main

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/go-grpc-example/proto/stream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

const (
	PORT = "9002"
)

func main() {

	//c, err := credentials.NewServerTLSFromFile("./conf/server.pem", "./conf/server.key")
	//if err != nil {
	//	log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	//}

	cert, err := tls.LoadX509KeyPair("./conf/server.pem", "./conf/server.key")
	if err != nil {
		log.Fatal("证书读取错误", err)
	}
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./conf/ca.crt")
	if err != nil {
		log.Fatal("ca证书读取错误", err)
	}
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	c := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ClientAuth: tls.RequireAndVerifyClientCert,
		// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})

	server := grpc.NewServer(grpc.Creds(c))

	streamObject := new(stream.StreamService)
	stream.RegisterStreamServiceServer(server, streamObject)

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
