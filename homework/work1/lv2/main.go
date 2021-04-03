package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	for{
		fmt.Println("请输入计算式")
		fmt.Scanln(&s)
		a,_:=strconv.Atoi(s[0:1])
		b,_:=strconv.Atoi(s[2:])
		c:=s[1:2]
		switch c{
		case "+"://加
			fmt.Println(add(a,b))
		case "-"://减
			fmt.Println(subtract(a,b))
		case "*"://乘
			fmt.Println(multiply(a,b))
		case "/"://除
			fmt.Println(divide(a,b))
		}
	}
}
//加
func add(a,b int)int{
	return a+b
}
//减
func subtract(a,b int)int{
	return a-b
}
//乘
func multiply(a,b int)int{
	return a*b
}
//除
func divide(a,b int)int{
	return a/b
}
