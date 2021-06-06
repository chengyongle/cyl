package controller//实现路由接口
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"wuziqi/game"
	"wuziqi/service"
)


func Register(c *gin.Context) {
	res:=service.Register(c)
	switch res{
	case 0:
		c.JSON(200,gin.H{
			"code":20000,
			"message":"账号已存在，请登录",
		})
	case 1:
		c.String(200,"注册成功")
		//c.JSON(200,gin.H{
		//	"code":10000,
		//	"message":"注册成功",
		//})
	case 2:
		c.JSON(200,gin.H{
			"code":    20000,
			"message": "注册失败",
		})
	}
}
func Login(c *gin.Context) {
	res:=service.Login(c)
	switch res{
	case 0:c.JSON(200,gin.H{
		"code":20000,
		"message":"账号不存在，请注册",
	})
	case 1:c.JSON(200,gin.H{
		"code":20000,
		"message":"登陆异常",
	})
	case 2:c.JSON(200,gin.H{
		"code":20000,
		"message":"密码错误",
	})
	case 3:
		//name:=c.PostForm("username")
		//c.SetCookie("user_cookie", name,
		//	1000, "/", "localhost", false, true)
		c.Writer.Write([]byte("登陆成功"))
	}
}
//加入房间
func Join(c *gin.Context){

	username:=c.Query("u")
	roomname:=c.Query("r")

	//R,ok:=game.Manager1.Rooms[roomname]
	//if !ok{
	//	c.String(200,"房间不存在")
	//	return
	//}
	//if R.Num_of_user==2{
	//	c.String(200,"房间已满")
	//	return
	//}
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

	client:=game.Client{
		Username: username,
		Ch:    nil,
		Roomname: roomname,
		Conn:     conn,
		Isready: false,
	}
	var room game.Room
	room.Jiaru(client)
	go client.Recv(&room)
	loop:
	for{
		select{
		case m:=<-client.Ch:
			if m=="1"{
				client.Isready=true
			}else if m=="0"{
				client.Isready=false
			}else if m=="q"{
				break loop
			}
		default:if room.Canstart(){
			room.Play(client.Conn)
		}
		}
	}
	close(client.Ch)
	client.Conn.Close()
	return
}
//创建房间
func Create(c *gin.Context){
	username:=c.Query("u")
	roomname:=c.Query("r")
	if _,ok:=game.Manager1.Rooms[roomname];ok{
		c.String(200,"房间已存在")
		return
	}
	upgrader := websocket.Upgrader{
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		c.String(200,"连接异常")
		return
	}
	client:=game.Client{
		Username: username,
		Ch:    nil,
		Roomname: roomname,
		Conn:     conn,
		Isready: false,
	}
	var room game.Room
	room.Chuangjian(client)
	loop:
	for{
		select{
		case m:=<-client.Ch:
			if m=="1"{
				fmt.Println("1")
				client.Isready=true
			}else if m=="0"{
				client.Isready=false
			}else if m=="q"{
				break loop
			}
			default:if room.Canstart(){
				room.Play(client.Conn)
			}
		}
	}
	close(client.Ch)
	client.Conn.Close()
	return
}