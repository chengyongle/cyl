package main

import (
	"fmt"
)

func main() {
	var client Client
	fmt.Println("登录还是注册？输入1登录，2注册")
	var a int
	for fmt.Scanln(&a);a!=1&&a!=2;fmt.Scanln(&a){
		fmt.Println("请重新输入，输入1登录，2注册")
	}
	switch a {
	case 1:
		for s:=login(&client);s!="登陆成功";s=login(&client){
			fmt.Println(s)
			continue
		}
		fmt.Println("登陆成功")
		break
	case 2:
		for s:=register(&client);s!="注册成功";s=register(&client){
			fmt.Println(s)
			continue
		}
		fmt.Println("注册成功")
		break
	}
	fmt.Println("输入1创建房间，输入2加入房间")
	for fmt.Scanln(&a);a!=1&&a!=2;fmt.Scanln(&a){
		fmt.Println("请重新输入，输入1创建房间，输入2加入房间")
	}
	switch a{
	case 1:createroom(&client)
	case 2:joinroom(&client)
	}
}
