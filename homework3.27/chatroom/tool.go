package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
//消息结构
type Message struct {
	User    string `json:"user"`
	Content string `json:"content"`
}
//所有房间
type Manager struct{
	rooms map[string]map[string]chan Message
	//key为房间名，value为用户名与信息管道的映射
	}
//用户结构
type Client struct{
	Username string
	Mesch chan Message
	Roomname string  //所在房间名
	Conn *websocket.Conn
}
//广播
func (m *Manager)Board (roomname string,mes Message){
	for _,ch:=range m.rooms[roomname]{
		ch<-mes
	}
}
//加入房间
func (m *Manager)Join(c *Client){
	if m.rooms[c.Roomname]==nil{
		m.rooms[c.Roomname]=map[string]chan Message{
			c.Username:c.Mesch,
		}
	}else{
		m.rooms[c.Roomname][c.Username]=c.Mesch
	}
}
//退出房间
func (m *Manager)Quit(c *Client){
	close(m.rooms[c.Roomname][c.Username])
	m.rooms[c.Roomname][c.Username]=nil
	s:="System message:"+c.Username+" exit"
	err:=c.Conn.WriteMessage(websocket.BinaryMessage,[]byte(s))
	if err!=nil{
		log.Println(err)
	}
	c.Conn.Close()
}
func (c *Client)Recv(){
	for{
		_,bytes,err:=c.Conn.ReadMessage()
		if err!=nil{
			log.Println(err)
			break
		}
		content:=string(bytes)
		mes:=Message{
			User:c.Username,
			Content: content,
		}
		manager.Board(c.Roomname,mes)
		if content=="/quit"{
			manager.Quit(c)
			return
		}
	}
}
func (c *Client)Write(){
	for{
		select{
		case m:=<-c.Mesch:
			if m.Content=="/quit"{
				return
			}
			s:=m.User+":"+m.Content
			err:=c.Conn.WriteMessage(websocket.BinaryMessage,[]byte(s))
			if err!=nil{
				log.Println(err)
				return
			}
		}
	}
}
var manager Manager
//用户请求里携带两个参数，用户名和房间名，若房间名不存在，则创建房间
func ws(c *gin.Context){
	upgrader := websocket.Upgrader{
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"status": 10022,
			"info": "failed",
		})
		return
	}
	uname:=c.Query("u")
	rname:=c.Query("r")
	client:=Client{
		Username: uname,
		Mesch:    make(chan Message),
		Roomname: rname,
		Conn:     conn,
	}
	manager.Join(&client)
	go client.Recv()
	go client.Write()
}

