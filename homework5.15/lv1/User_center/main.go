//server（被调用rpc的一方）
package main

import (
	"google.golang.org/grpc"
	pb "homework5.15/lv1/User_center/pbfile/proto"
	"log"
	"net"
	"homework5.15/lv1/User_center/dao"
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