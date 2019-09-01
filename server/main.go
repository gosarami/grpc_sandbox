package main

import (
	"context"
	"log"
	"net"
	pb "grpc_sandbox/protobuf"
	"google.golang.org/grpc"
)

func main() {
	listener,err := net.Listen("tcp", ":5300")
	if err != nil {
		log.Fatalf("failed to listen %v¥n",err)
		return
	}
	grpcSrv := grpc.NewServer()
	pb.RegisterGreeterServer(grpcSrv,&server{})
	log.Printf("hello service is running!")
	grpcSrv.Serve(listener)
}

type server struct{}

func (s *server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error){
	return &pb.HelloReply{
		Message: "hello! " + request.GetName(),
	},nil
}

