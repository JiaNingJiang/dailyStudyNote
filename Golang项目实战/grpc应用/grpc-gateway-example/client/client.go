package main

import (
	"context"
	"github.com/grpc-gateway-example/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//creds, err := credentials.NewClientTLSFromFile("./cert/server.pem", "www.github.com")
	//if err != nil {
	//	log.Println("Failed to create TLS credentials %v", err)
	//}
	//conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	c := proto.NewHelloWorldClient(conn)
	context := context.Background()
	body := &proto.HelloWorldRequest{
		Referer: "Grpc",
	}

	r, err := c.SayHelloWorld(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Message)
}
