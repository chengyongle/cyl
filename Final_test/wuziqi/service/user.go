package service
import (
	"github.com/gin-gonic/gin"
	"wuziqi/models"
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