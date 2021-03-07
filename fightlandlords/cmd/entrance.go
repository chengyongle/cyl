package cmd
import (
	"github.com/gin-gonic/gin"
	"fightlandlords/controller"
)
//定义路由接口
func Entrance() {
	e:=gin.Default()
	e.POST("/register",controller.Register)
	e.POST("/login",controller.Login)
	e.POST("/exit",controller.Exit)

	e.Run()
}
