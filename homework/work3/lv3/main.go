package main

import "fmt"

//func main() {
//	over := make(chan bool)
//	for i := 0; i < 10; i++ {
//		go func() {
//			fmt.Println(i)
//		}()//闭包，此处传入的是i的引用，最后始终调用同一i值
//		if i == 9 {
//			over <- true
//		}
//	}
//	<-over//死锁
//	fmt.Println("over!!!")
//}
func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			over<-true
		}(i)
		<-over
	}
	fmt.Println("over!!!")
	demo2()
}
//也可以这样
func demo2() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			over<-true
		}(i)
	}
	for i:=0;i<10;i++{
		<-over
	}
	fmt.Println("over!!!")
}