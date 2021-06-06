package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"net/url"
)
type Client struct{
	Username string
	Password string
	Roomname string  //房间名
	Conn *websocket.Conn
}

var serverch=make(chan string)
var sendch=make(chan string)
var reader = bufio.NewReader(os.Stdin)
func (c *Client)Write(){
	for{
		line,err:=reader.ReadString('\n')
		if err!= nil {
			log.Println(err)
		}
		line =strings.Trim(line, "\n")
		sendch<-line
	}
}
func (c *Client)Recv(){
	for{
		_,bytes,err:=c.Conn.ReadMessage()
		if err!=nil{
			log.Println(err)
			break
		}
		serverch<-string(bytes)
	}
}
func register(c *Client)string{
	fmt.Println("请输入你的用户名")
	fmt.Scanln(&c.Username)
	fmt.Println("请输入你的密码")
	fmt.Scanln(&c.Password)
	u:="http://127.0.0.1:8080/register"
	data:=url.Values{"username":{c.Username},"password":{c.Password}}
	resp,err:=http.PostForm(u,data)
	if err!=nil{
		return "注册异常"
	}
	defer resp.Body.Close()
	body, err:= ioutil.ReadAll(resp.Body)
	if err != nil {
		return "注册异常"
	}
	return string(body)
}
func login(c *Client)string{
	fmt.Println("请输入你的用户名")
	fmt.Scanln(&c.Username)
	fmt.Println("请输入你的密码")
	fmt.Scanln(&c.Password)
	u:="http://127.0.0.1:8080/login"
	data:=url.Values{"username":{c.Username},"password":{c.Password}}
	resp,err:=http.PostForm(u,data)
	if err!=nil{
		return "登陆异常"
	}
	defer resp.Body.Close()
	body, err:= ioutil.ReadAll(resp.Body)
	if err != nil {
		return "登陆异常"
	}
	return string(body)
}
func createroom(c *Client){
	fmt.Println("请输入你要创建的房间名")
	var rname string
	fmt.Scanln(&rname)
	c.Roomname=rname
	str:="ws://127.0.0.1:8080/play/create?u="+c.Username+"&r="+c.Roomname
	conn,rsp,err:=websocket.DefaultDialer.Dial(str,nil)
	if err!=nil{
		fmt.Println(err)
	}
	body,_:= ioutil.ReadAll(rsp.Body)
	if string(body)!=""{
		fmt.Println(string(body))
		fmt.Println("mark")
		return
	}
	c.Conn=conn
	fmt.Println("创建房间成功，输入1准备")
	var p string
	for fmt.Scanln(&p);p!="1";fmt.Scanln(&p){}
	err=conn.WriteMessage(websocket.BinaryMessage,[]byte(p))
	for{
		_,bytes,err:=c.Conn.ReadMessage()
		if err!=nil{
			log.Println(err)
			return
		}
		if string(bytes)=="game start"{
			break
		}
	}
	blackplay(conn)
}

func joinroom(c *Client){
	fmt.Println("请输入你要加入的房间名")
	var rname string
	fmt.Scanln(&rname)
	c.Roomname=rname
	str:="ws://127.0.0.1:8080/play/join?u="+c.Username+"&r="+c.Roomname
	conn,rsp,err:=websocket.DefaultDialer.Dial(str,nil)
	if err!=nil{
		log.Println(err)
	}
	body,_:= ioutil.ReadAll(rsp.Body)
	if string(body)!=""{
		fmt.Println(string(body))
		fmt.Println("mark")
		return
	}
	c.Conn=conn
	fmt.Println("加入房间成功，输入1准备")
	var p string

	for fmt.Scanln(&p);p!="1";fmt.Scanln(&p){}
	err=conn.WriteMessage(websocket.BinaryMessage,[]byte(p))
	for{
		_,bytes,err:=c.Conn.ReadMessage()
		if err!=nil{
			log.Println(err)
			return
		}
		if string(bytes)=="game start"{
			break
		}
	}
	whiteplay(conn)

}
func whiteplay(conn *websocket.Conn){
	var step string
	for{
		_,bytes,_:=conn.ReadMessage()
		fmt.Println(string(bytes))
		if _,res,err:=conn.ReadMessage();err==nil{
			fmt.Println(string(res))
			return
		}
		for{
			fmt.Println("输入你的落子，两位数，第一位为横轴，第二位为纵轴")
			fmt.Scanln(&step)
			a,_:=strconv.Atoi(step)
			if a>100||a<0{
				fmt.Println("输入不合法，请重新输入")
				continue
			}
			break
		}
		s:="w"+step
		conn.WriteMessage(websocket.BinaryMessage,[]byte(s))
	}
}
func blackplay(conn *websocket.Conn){
	var step string
	for{
		for{
			fmt.Println("输入你的落子，两位数，第一位为横轴，第二位为纵轴")
			fmt.Scanln(&step)
			a,_:=strconv.Atoi(step)
			if a>100||a<0{
				fmt.Println("输入不合法，请重新输入")
				continue
			}
			break
		}
		s:="b"+step
		conn.WriteMessage(websocket.BinaryMessage,[]byte(s))
		_,bytes,_:=conn.ReadMessage()
		fmt.Println(string(bytes))
		if _,res,err:=conn.ReadMessage();err==nil{
			fmt.Println(string(res))
			return
		}
	}
}