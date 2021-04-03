package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
//非法字符
var char1="!@#$%^&*()_+:,./`<>?{}"
var char2="!@#$%^&*()_+:,./`<>?{}qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
var m=make(map[string]string)
func main() {
	f,err:=os.OpenFile("./users.data",os.O_CREATE|os.O_RDWR,0644)
	if err!=nil{
		fmt.Printf("open file err:%v",err)
		return
	}
	defer f.Close()
	var bytes []byte
	bytes,err=ioutil.ReadAll(f)
	os.Truncate("./users.data",0)
	f.Seek(0,0)
	if err!=nil{
		fmt.Printf("read err:%v",err)
		return
	}
	json.Unmarshal(bytes,&m)
	for{
		fmt.Println("请输入你要进行的操作序号")
		fmt.Println("1.登录")
		fmt.Println("2.注册")
		fmt.Println("3.退出")
		var a int
		fmt.Scanln(&a)
		switch a{
		case 1: login()
		case 2: register()
		case 3:	exit(f)
				return
		}
	}
}
func login(){
	fmt.Println("请输入账号")
	var name string
	fmt.Scanln(&name)
	fmt.Println("请输入密码")
	var password string
	fmt.Scanln(&password)
	if m[name]==""{
		fmt.Println("该账号不存在")
		return
	}
	if m[name]!=password{
		fmt.Println("密码错误")
		return
	}
	fmt.Println("登录成功！")
}
func register(){
	var name,password string
	for{
		fmt.Println("请输入账号")
		fmt.Scanln(&name)
		if len(name)<8{
			fmt.Println("账号不能少于8位")
			continue
		}
		if strings.ContainsAny(name,char2){
			fmt.Println("账号只能是数字")
			continue
		}
		if m[name]!=""{
			fmt.Println("该账号已存在")
			continue
		}
		break
	}
	for{
		fmt.Println("请输入密码")
		fmt.Scanln(&password)
		if len(name)<6{
			fmt.Println("密码不能少于6位")
			continue
		}
		if strings.ContainsAny(name,char1){
			fmt.Println("密码由数字和字母组成，不能包含其他符号")
			continue
		}
		break
	}
	m[name]=password
	fmt.Println("注册成功！")
}
func exit(f *os.File){
	b,err:=json.Marshal(m)
	if err!=nil{
		fmt.Printf("err%v",err)
		return
	}
	f.WriteString(string(b))
	fmt.Println("成功退出")
}