package  controller
//实现路由接口
import (
	"github.com/gin-gonic/gin"
	"awesomeProject/service"
)


func Register(c *gin.Context) {
	res:=service.Register(c)
	switch res{
	case 0:c.JSON(200,gin.H{
		"code":20000,
		"message":"账号已存在，请登录",
	})
	case 1:c.JSON(200,gin.H{
		"code":20000,
		"message":"注册失败",
	})
	case 2:c.JSON(200,gin.H{
		"code":10000,
		"message":"注册成功",
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
		name:=c.PostForm("username")
		c.SetCookie("user_cookie", name,
		1000, "/", "localhost", false, true)
		c.JSON(200,gin.H{
		"code":10000,
		"message":"登陆成功",
		"你好":name,
	})
	}
}
func Exit(c *gin.Context){
	c.SetCookie("user_cookie", "",
		-1, "/", "localhost", false, true)
	c.JSON(200,gin.H{
		"code":10000,
		"message":"已注销",
	})
}
func SendMessage(c *gin.Context) {
	res:=service.SendMessage(c)
	switch res{
	case 0:c.JSON(200,gin.H{
		"code":20000,
		"message":"请先登录",
	})
	case 1:c.JSON(200,gin.H{
		"code":20000,
		"message":"评论失败",
	})
	case 2:c.JSON(200,gin.H{
		"code":10000,
		"message":"评论成功",
	})
	}
}
func ViewMessage(c *gin.Context){
	res:=service.ViewMessage(c)
	if res==""{
		c.JSON(200,gin.H{
			"code":20000,
			"message":c.PostForm("name")+"未留言过",
		})
	}else{
		c.JSON(200,gin.H{
			"code":10000,
			"message":res,
		})
	}
}
func ReplyMessage(c *gin.Context){
	res:=service.ReplyMessage(c)
	switch res{
	case 0:c.JSON(200,gin.H{
		"code":"20000",
		"message":"请先登录",
	})
	case 1:c.JSON(200,gin.H{
		"code":"20000",
		"message":"回复异常",
	})
	case 2:c.JSON(200,gin.H{
		"code":"10000",
		"message":"回复成功",
	})
	}
}
func Interesting(c *gin.Context){
	res:=service.Interesting(c)
	switch res{
	case 0:c.JSON(200,gin.H{
		"code":"20000",
		"message":"请先登录",
	})
	case 1:c.JSON(200,gin.H{
		"code":"20000",
		"message":"点赞失败",
	})
	case 2:c.JSON(200,gin.H{
		"code":"10000",
		"message":"点赞成功",
	})
	}
}