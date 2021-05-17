package cmd

import (
	"fightlandlords/controller"
	"fightlandlords/service"
	"fmt"
	"github.com/gin-gonic/gin"
)
//定义路由接口
func Entrance() {
	e:=gin.Default()
	e.POST("/register",controller.Register)
	e.POST("/login",controller.JWTmiddleware(),controller.Login)
	e.POST("/exit",controller.Exit)
	e.POST("/update",controller.Update)
	e.GET("/test",test)
	e.Run()
}
//测试
func test(c *gin.Context){
	authHeader:=c.Request.Header.Get("Authorization")
	fmt.Println(authHeader)
	jwt,err:=service.Check(authHeader)
	if err!=nil{
		fmt.Printf("err:%v",err)
		return
	}
	c.JSON(200,gin.H{
		"code":20000,
		"message":"你好，"+jwt.Payload.Username,
	})
}
