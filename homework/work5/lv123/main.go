package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	e:=gin.Default()
	e.Use(Middleare())
	e.POST("/login",login)
	e.POST("/register",register)
	e.POST("/upload",upload)
	e.POST("/draw",draw)
	e.Run()
}
