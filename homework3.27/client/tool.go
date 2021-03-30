package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"net/url"
)
type Client struct{
	Username string
	Password string
	Roomname string  //房间名
	Conn *websocket.Conn
}
type Message struct{
	Name string
	Content string
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
