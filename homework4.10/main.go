package main

import (
	"fightlandlords/cmd"
	"fightlandlords/dao"
)

func main() {
	dao.SqlInit()
	cmd.Entrance()
}
