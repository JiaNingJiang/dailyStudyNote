package main

import (
	"context"
	"github.com/go-grpc-example/pkg/gtls"
	"github.com/go-grpc-example/proto/search"
	"google.golang.org/grpc"
	"log"
	"time"
)

const PORT = "9003"

func main() {

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(5*time.Second))) // 设置定时context
	defer cancel()

	tlsClient := gtls.Client{
		ServerName: "go-grpc-example.github.com",
		CertFile:   "./conf/server.pem",
	}
	c, err := tlsClient.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}

	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := search.NewSearchServiceClient(conn)
	resp, err := client.Search(ctx, &search.SearchRequest{ // search服务不再使用context.Background()，而是使用新的定时ctx
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}
