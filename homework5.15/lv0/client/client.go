//client(调用rpc的一方)
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "homework5.15/lv0/client/proto"
	"log"
)

const (
	address     = "localhost:50051"
)

func main() {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDengluClient(conn)


	for  {
		//这段不重要
		fmt.Println("input username&password:")
		iptName := ""
		_, _ = fmt.Scanln(&iptName)
		iptPassword := ""
		_, _ = fmt.Scanln(&iptPassword)

		loginResp, _ := c.Login(context.Background(), &pb.LoginReq{
			UserName: iptName,
			PassWord: iptPassword,
		})

		if loginResp.OK{
			fmt.Println("success")
			break
		}
		fmt.Println("retry")
	}
}
