package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"awesomeProject/models"
	"strconv"
)
func SendMessage(c *gin.Context)int{
	username,err:=c.Cookie("user_cookie")
	if err!=nil{
		//未登录
		return 0
	}
	fmt.Println(username)
	message:=c.PostForm("message")
	res:=models.SaveMessage(username,message)
	return res
}
func ViewMessage(c *gin.Context)string{
	name:=c.PostForm("name")
	messages:=models.ViewMessage(name)
	bytes,_:=json.Marshal(messages)
	return string(bytes)
}
func ReplyMessage(c *gin.Context)int{
	username,err:=c.Cookie("user_cookie")
	if err!=nil{
		//未登录
		return 0
	}
	fmt.Println(username)
	id,err:=strconv.Atoi(c.PostForm("id"))
	if err!=nil{
		fmt.Printf("failed:%v",err)
		return 1
	}
	message:=c.PostForm("message")
	res:=models.ReplyMessage(username,message,id)
	return res
}
func Interesting(c *gin.Context)int{
	_,err:=c.Cookie("user_cookie")
	if err!=nil{
		//未登录
		return 0
	}
	id,err:=strconv.Atoi(c.PostForm("id"))
	if err!=nil{
		fmt.Printf("failed:%v",err)
		return 1
	}
	res:=models.Interesting(id)
	return res
}