package main

import (
	"fmt"
)

var (
	ch=make(chan int)
	myres = make(map[int]int, 20)
)

func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myres[n] = res
	<-ch
}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
		ch<-1
	}
	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
}