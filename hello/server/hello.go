package main

import (
	"context"
	"log"

	pb "grpc-unary/hello/proto"
)

func (s *Server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Greet was invoked with %v\n", in)
	return &pb.HelloResponse{Result: "Hello " + in.FirstName}, nil
}
