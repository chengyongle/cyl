package main

import (
	"fmt"
	"time"
)
var ch=make(chan int)
func main() {
	go qifei()
	go dingshiqi(2,"没有困难的工作，只有勇敢的打工人！")
	go dingshiqi(8,"早安，打工人！")
	<-ch
}
func qifei(){
	ticker:=time.NewTicker(time.Hour)
	for{
		select{
		case <-ticker.C:
			fmt.Println("芜湖！起飞！")
		}
	}
}
//定时器
func dingshiqi(hour int,s string){
	nowtime:=time.Now()
	t:=time.Date(nowtime.Year(),nowtime.Month(),
		nowtime.Day(),hour,0,0,0,time.Local)
	if t.Before(nowtime) {
		t = t.Add(time.Hour*24)
	}
	timer:=time.NewTimer(t.Sub(nowtime))
	for{
		<-timer.C
		fmt.Println(s)
		timer.Reset(time.Hour*24)
	}
}
