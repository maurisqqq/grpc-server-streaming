package main

import (
	"context"
	"io"
	"log"

	pb "grpc-server-streaming/hello/proto"
)

func doHelloManyTimes(c pb.HelloServiceClient) {
	log.Println("doHelloManyTimes was invoked")

	req := &pb.HelloRequest{
		FirstName: "Maulana",
	}

	stream, err := c.HelloManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling HelloManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("HelloManyTimes: %s\n", msg.Result)
	}
}
