package service

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
)
func Register(c *gin.Context)int{
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	res:=models.Register(username,password)
	return res
}
func Login(c *gin.Context)int{
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	res:=models.Login(username,password)
	return res
}