package main

import (
	"chatroom/dao"
	"github.com/gin-gonic/gin"
)
func main(){
	dao.SqlInit()
	manager=Manager{
		rooms: make(map[string]map[string]chan Message),
	}
	e:=gin.Default()
	e.POST("/register",dao.Register)
	e.POST("/login",dao.Login)
	e.GET("/ws",ws)
	e.Run("127.0.0.1:8080")
}

