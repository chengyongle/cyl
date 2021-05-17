package main

import (
	"context"
	"google.golang.org/grpc"
	pb "homework5.15/lv0/server/pbfile/proto"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{} //服务对象

func (s server) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp,error) {
	resp := &pb.LoginResp{}
	log.Println("recv:",req.UserName,req.PassWord)
	if req.PassWord != GetPassWord(req.UserName){
		resp.OK = false
		return resp,nil
	}
	resp.OK = true
	return resp,nil
}

func GetPassWord(userName string)(password string)    {
	return userName+"123456"
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDengluServer(s,&server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}