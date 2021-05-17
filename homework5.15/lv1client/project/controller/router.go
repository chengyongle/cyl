package  controller
//实现路由接口
import (
	"fightlandlords/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(200,gin.H{
		"code":10000,
		"message":"注册成功",
	})
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
		name:=c.PostForm("username")
		//c.SetCookie("user_cookie", name,
		//	1000, "/", "localhost", false, true)
		token:=service.NewJWT(name)
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "登录成功",
			"token": token.Token,
		})
	}
}
func Update(c *gin.Context){
	res:=service.Update(c)
	if res{
		c.JSON(200,gin.H{
			"code":10000,
			"message":"修改密码成功",
		})
	}else{
		c.JSON(200,gin.H{
			"code":10000,
			"message":"请先登录",
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
func JWTmiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader:=c.Request.Header.Get("Authorization")
		jwt,err:=service.Check(authHeader)
		if err!=nil{
			fmt.Printf("err:%v",err)
			return
		}
		c.JSON(200,gin.H{
			"code":20000,
			"message":"你好，"+jwt.Payload.Username,
		})
		c.Next()
	}
}