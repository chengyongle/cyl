package cmd
import (
	"github.com/gin-gonic/gin"
	"awesomeProject/controller"
)
//定义路由接口
func Entrance() {
	e:=gin.Default()
	e.POST("/register",controller.Register)
	e.POST("/login",controller.Login)
	e.POST("/sendm",controller.SendMessage)
	e.POST("/viewm",controller.ViewMessage)
	e.POST("/replym",controller.ReplyMessage)
	e.POST("/exit",controller.Exit)
	e.POST("/interesting",controller.Interesting)
	e.Run()
}
