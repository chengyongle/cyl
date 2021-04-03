package main

import (
	"awesomeProject/cmd"
	"awesomeProject/dao"
)

func main() {
	dao.SqlInit()
	cmd.Entrance()
}