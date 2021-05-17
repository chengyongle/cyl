package  models

import (
	"context"
	pb "fightlandlords/pbfile/proto"
	"google.golang.org/grpc"
	"log"
)
const (
	address     = "localhost:50051"
)
func Register(name string,password string)int32{
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	C := pb.NewUserserverClient(conn)
	registerResp, _ := C.Register(context.Background(), &pb.RegisterReq{
		UserName: name,
		PassWord: password,
	})
	return registerResp.RES
}
func Login(name string,password string)int32{
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	C := pb.NewUserserverClient(conn)
	loginResp, _ := C.Login(context.Background(), &pb.LoginReq{
		UserName: name,
		PassWord: password,
	})
	return loginResp.RES

}
func Update(name,newpassword string)bool{
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	C := pb.NewUserserverClient(conn)
	updateResp, _ := C.Update(context.Background(), &pb.UpdateReq{
		UserName: name,
		NewPassWord: newpassword,
	})
	return updateResp.OK
}
