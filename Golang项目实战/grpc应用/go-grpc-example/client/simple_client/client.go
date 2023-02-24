package main

import (
	"context"
	"github.com/go-grpc-example/client/auth"
	"github.com/go-grpc-example/proto/search"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"google.golang.org/grpc"
)

const PORT = "9001"

func main() {

	user := &auth.Authentication{
		User:     "admin",
		Password: "admin",
	}

	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(user))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := search.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &search.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}
