package main

import (
	"fmt"
	"log"

	pb "grpc-server-streaming/hello/proto"
)

func (*Server) HelloManyTimes(in *pb.HelloRequest, stream pb.HelloService_HelloManyTimesServer) error {
	log.Printf("HelloManyTimes was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.HelloResponse{
			Result: res,
		})
	}

	return nil
}
