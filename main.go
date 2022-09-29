package main

import (
	"log"
	"net"

	pb "grpc-unary/book"
	"grpc-unary/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error on listen %v", err.Error())
	}

	s := grpc.NewServer()

	// register
	pb.RegisterBookServiceServer(s, &server.Server{})
	reflection.Register(s)
	
	s.Serve(listen)
}
