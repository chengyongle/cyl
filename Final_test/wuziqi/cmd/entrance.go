package cmd

import (
	"github.com/gin-gonic/gin"
	"wuziqi/controller"
)
//定义路由接口
func Entrance() {

	e:=gin.Default()
	e.POST("/register",controller.Register)
	e.POST("/login",controller.Login)
	playgroup:=e.Group("/play"){
		playgroup.GET("/create",controller.Create)
		playgroup.GET("/join",controller.Join)
	}
	e.Run()
}