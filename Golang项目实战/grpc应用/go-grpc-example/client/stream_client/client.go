package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/go-grpc-example/proto/stream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"io/ioutil"
	"log"
)

const (
	PORT = "9002"
)

func main() {

	//c, err := credentials.NewClientTLSFromFile("./conf/server.pem", "www.github.com")
	//if err != nil {
	//	log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	//}

	cert, _ := tls.LoadX509KeyPair("./conf/client.pem", "./conf/client.key")
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("./conf/ca.crt")
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	c := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ServerName: "*.github.com",
		RootCAs:    certPool,
	})

	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()

	client := stream.NewStreamServiceClient(conn)

	err = printLists(client, &stream.StreamRequest{Pt: &stream.StreamPoint{Name: "gRPC Stream Client: List", Value: 2018}})
	if err != nil {
		log.Fatalf("printLists.err: %v", err)
	}

	err = printRecord(client, &stream.StreamRequest{Pt: &stream.StreamPoint{Name: "gRPC Stream Client: Record", Value: 2018}})
	if err != nil {
		log.Fatalf("printRecord.err: %v", err)
	}

	//err = printRoute(client, &stream.StreamRequest{Pt: &stream.StreamPoint{Name: "gRPC Stream Client: Route", Value: 2018}})
	//if err != nil {
	//	log.Fatalf("printRoute.err: %v", err)
	//}
}

func printLists(client stream.StreamServiceClient, r *stream.StreamRequest) error {
	stream, err := client.List(context.Background(), r) // 发送请求的同时形成stream流对象
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv() // 利用流对象接收服务端回复
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}

	return nil
}

func printRecord(client stream.StreamServiceClient, r *stream.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)

	return nil
}

func printRoute(client stream.StreamServiceClient, r *stream.StreamRequest) error {
	stream, err := client.Route(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n <= 6; n++ {
		err = stream.Send(r) // 发送一次
		if err != nil {
			return err
		}

		resp, err := stream.Recv() // 接收一次
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}

	stream.CloseSend()

	return nil
}
