package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"user_center/dao"
	pb "user_center/pbfile/proto"
)

const (
	port = ":50051"
)

func main() {
	dao.SqlInit()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //起一个服务
	pb.RegisterUserserverServer(s,&server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
