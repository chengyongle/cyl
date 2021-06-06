package main
import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"user_center/dao"
	pb "user_center/pbfile/proto"
	"log"
)

type server struct{} //服务对象

func (s server) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp,error) {
	resp := &pb.LoginResp{}
	log.Println("recv:",req.UserName,req.PassWord)
	resp.RES=denglu(req.UserName,req.PassWord)
	return resp,nil
}

func denglu(userName ,password string)int32{
	var user dao.User
	err:=dao.DB.Where("username=?",userName).Find(&user).Error
	if err==nil{
		//账号存在
		if user.Password==password{
			//登陆成功
			return 3
		}else{
			//密码错误
			return 2
		}
	}else if err==gorm.ErrRecordNotFound{
		//账号不存在
		return 0
	}else{
		//查询异常
		fmt.Printf("err:%v",err)
		return 1
	}
}
func (s server)Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp,error){
	resp := &pb.RegisterResp{}
	log.Println("recv:",req.UserName,req.PassWord)
	resp.RES=Isexist(req.UserName,req.PassWord)
	return resp,nil
}
func Isexist(username ,password string)int32{
	var user dao.User
	err:=dao.DB.Where("username=?",username).Find(&user).Error
	if err==nil{
		//账号已存在
		return 0
	}else if err==gorm.ErrRecordNotFound{
		//账号不存在
		user.Username=username
		user.Password=password
		dao.DB.Create(&user)
		return 1
	}else{
		//其他错误
		fmt.Printf("err:%v",err)
		return 2
	}
}