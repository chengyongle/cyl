package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f,err:=os.Create("./proverb.txt")
	if err!=nil{
		fmt.Printf("err:%v",err)
	}
	defer f.Close()
	s:=[]byte("don't communicate by sharing memory share memory by communicating")
	_,err=f.Write(s)
	if err!=nil{
		fmt.Printf("err:%v",err)
	}
	//var b=make([]byte,50)
	//_,err=f.Read(b)
	//if err!=nil{
	//	fmt.Printf("err:%v",err)
	//}
	b,_:=ioutil.ReadFile("./proverb.txt")
	fmt.Println(string(b))
}
