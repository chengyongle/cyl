package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func main(){
	var client Client
	fmt.Println("登录还是注册？输入1登录，2注册")
	var a int
	fmt.Scanln(&a)
	SWITCH:
	for{
		switch a {
			case 1:
				for s:=login(&client);s!="登陆成功";s=login(&client){
					fmt.Println(s)
					continue
				}
				fmt.Println("登陆成功")
				break SWITCH
			case 2:
				for s:=register(&client);s!="注册成功";s=register(&client){
					fmt.Println(s)
					continue
				}
				fmt.Println("注册成功")
				break SWITCH
			default:
				fmt.Println("请重新输入，输入1登录，2注册")
				fmt.Scanln(&a)
				continue
		}
	}
	fmt.Println("请输入你要加入的房间名（房间不存在则直接创建）")
	fmt.Scanln(&client.Roomname)
	str:="ws://127.0.0.1:8080/ws?u="+client.Username+"&r="+client.Roomname
	conn,_,err:=websocket.DefaultDialer.Dial(str,nil)
	if err!=nil{
		log.Println(err)
	}
	fmt.Println("加入房间成功，开始聊天吧！")
	client.Conn=conn
	go client.Write()
	go client.Recv()
	for{
		select{
		case rmes:=<-serverch:
			fmt.Println(rmes)
		case smes:=<-sendch:
			err=conn.WriteMessage(websocket.BinaryMessage,[]byte(smes))
			if err != nil {
				log.Println(err)
			}
			if smes=="/quit"{
				fmt.Println("客户端退出")
				conn.Close()
				break
			}
		}
	}
}


