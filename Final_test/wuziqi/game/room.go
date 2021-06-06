package game
//消息结构
import (
	"github.com/gorilla/websocket"
	"log"
	"strconv"
)
//实例化大堂经理
var Manager1=Manager{
	Rooms: make(map[string]Room),
}
//大堂经理
type Manager struct{
	Rooms map[string]Room
}
//房间类
type Room struct {
	Roomname string
	Num_of_user int  //人数
	Num_of_ready int   //准备数
	clients []Client  //
	Chess [10][10]int  //棋局
}

//用户结构
type Client struct{
	Username string
	Ch chan string
	Roomname string  //所在房间名
	Conn *websocket.Conn
	Isready bool// 是否准备
}
//创建房间
func (r *Room)Chuangjian(c Client){
	r.Roomname=c.Roomname
	r.Num_of_user=1
	r.clients=append(r.clients,c)
}
//加入房间
func (r *Room)Jiaru(c Client){
	r.Roomname=c.Roomname
	r.Num_of_user++
	r.clients=append(r.clients,c)
}
func (r *Room)Canstart()bool{
	if r.clients[0].Isready==true&&r.clients[1].Isready==true{
		return true
	}else{
		return false
	}
}
func (c *Client)Recv(r *Room){
	for{
		for r.Canstart(){}
		_,bytes,err:=c.Conn.ReadMessage()
		if err!=nil{
			log.Println(err)
			break
		}
		content:=string(bytes)
		c.Ch<-content
	}
}
func (r *Room)Play(conn *websocket.Conn){
	s:="game start"
	err:=conn.WriteMessage(websocket.BinaryMessage,[]byte(s))
	if err!=nil{
		log.Println(err)
		return
	}
	for{
		_,bytes,err:=conn.ReadMessage()
		if err!=nil{
			log.Println(err)
			break
		}
		step:=string(bytes)
		color:=step[0:1]
		x,_:=strconv.Atoi(step[1:2])
		y,_:=strconv.Atoi(step[2:3])
		if color=="b"{
			r.Chess[x][y]=1
			s="black move x"+step[1:2]+"y"+step[2:3]
		}else{
			r.Chess[x][y]=2
			s="white move x"+step[1:2]+"y"+step[2:3]
		}
		conn.WriteMessage(websocket.BinaryMessage,[]byte(s))
		if r.Judge(x,y){
			var res string
			if color=="b"{
				res="black win!"
			}else{
				res="white win!"
			}
			conn.WriteMessage(websocket.BinaryMessage,[]byte(res))
			return
		}
	}
}
