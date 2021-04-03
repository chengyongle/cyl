package main

import (
	"fmt"
	"sync"
)
var ch1=make(chan int)
var ch2=make(chan int)
var wg sync.WaitGroup


func main() {
	wg.Add(2)
	go printj()
	go printo()
	ch1<-1
	wg.Wait()
}
//打印奇数
func printj(){
	for {
		v:=<-ch1
		fmt.Println(v)
		ch2<-v+1
		if v==99{
			break
		}
	}
	wg.Done()
}
//打印偶数
func printo(){
	for {
		v:=<-ch2
		fmt.Println(v)
		if v==100{
			break
		}
		ch1<-v+1
	}
	wg.Done()
}
