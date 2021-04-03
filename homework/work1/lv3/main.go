package main

import "fmt"

func main(){
	//a:=[]int{2,-8,7,-1,4,0}
	var b []int
	for{
		//终端输入数组，输入任意字母停止输入
		var num int
		_,err:=fmt.Scanln(&num)
		if err!=nil{
			break
		}
		b=append(b,num)
	}
	//fmt.Println(min(a))
	fmt.Println(min(b))
}
//找最小绝对值函数
func min(arr []int)int{
	res:=arr[0]//res保存最小的数
	p:=arr[0]*arr[0]//p保存其平方
	for _,i:=range arr{
		b:=i*i
		//平方最小的数就是绝对值最小的
		if b<p{
			p=b
			res=i
		}
	}
	return  res
}