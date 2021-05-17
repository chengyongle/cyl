package service

import (
	"fightlandlords/models"
	"github.com/gin-gonic/gin"
)
func Register(c *gin.Context)int32{
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	res:=models.Register(username,password)
	return res
}
func Login(c *gin.Context)int32{
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	res:=models.Login(username,password)
	return res
}
func Update(c *gin.Context)bool{
	authHeader:=c.Request.Header.Get("Authorization")
	if authHeader==""{
		return false
	}
	username:=c.PostForm("username")
	newpassword:=c.PostForm("newpassword")
	res:=models.Update(username,newpassword)
	return res
}
